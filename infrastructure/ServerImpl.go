package infrastructure

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	Openapi "github.com/nekochans/portfolio-backend/infrastructure/openapi"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type ServerImpl struct {
	httpHandler http.Handler
	router *chi.Mux
	DB     *sql.DB
	Logger *zap.Logger
}

type Members []*Openapi.Member

type GetMembersResponse struct {
	Items Members `json:"items"`
}

func(si *ServerImpl) GetMembers(w http.ResponseWriter, r *http.Request) {
	var ms Members

	ms = append(
		ms,
		&Openapi.Member{
			Id:             1,
			GithubUserName: "keitakn",
			GithubPicture:  "https://avatars1.githubusercontent.com/u/11032365?s=460&v=4",
			CvUrl:          "https://github.com/keitakn/cv",
		},
	)

	ms = append(
		ms,
		&Openapi.Member{
			Id:             2,
			GithubUserName: "kobayashi-m42",
			GithubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
			CvUrl:          "https://github.com/kobayashi-m42/cv",
		},
	)

	res := GetMembersResponse{Items: ms}
	CreateJsonResponse(w, r, http.StatusOK, res)
}

func(si *ServerImpl) GetMemberById(w http.ResponseWriter, r *http.Request) {
	si.httpHandler = Openapi.GetMembersCtx(si.httpHandler)

	id := r.Context().Value("id").(int)

	log.Println("🐱🐱")
	log.Println(id)
	log.Println("🐱🐱")

	res := &Openapi.Member{
		Id:             2,
		GithubUserName: "kobayashi-m42",
		GithubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
		CvUrl:          "https://github.com/kobayashi-m42/cv",
	}

	CreateJsonResponse(w, r, http.StatusOK, res)
}

func(si *ServerImpl) GetWebservices(w http.ResponseWriter, r *http.Request) {
}

func(si *ServerImpl) CreateJsonResponse(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
	res, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Fatal(err, "http.ResponseWriter() Fatal.")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Request-Id", middleware.GetReqID(r.Context()))
	w.WriteHeader(status)
	_, err = w.Write([]byte(res))
	if err != nil {
		log.Fatal(err, "http.ResponseWriter() Fatal.")
	}
}

func (si *ServerImpl) middleware() {
	si.router.Use(middleware.RequestID)
	si.router.Use(Logger(si.Logger))
	si.router.Use(middleware.Recoverer)
	si.router.Use(middleware.Timeout(time.Second * 60))
}

func (si *ServerImpl) Init(env string) {
	log.Printf("env: %s", env)
	si.middleware()
}

func NewServerImpl(logger *zap.Logger, router *chi.Mux) *ServerImpl {
	return &ServerImpl{
		router: router,
		Logger: logger,
	}
}

func NewServerImplWithMySQL(db *sql.DB, logger *zap.Logger, router *chi.Mux) *ServerImpl {
	return &ServerImpl{
		router: router,
		DB:     db,
		Logger: logger,
	}
}

func StartServerImpl() {
	var (
		port = flag.String("port", "8888", "addr to bind")
		env  = flag.String("env", "develop", "実行環境 (production, staging, develop)")
	)
	flag.Parse()

	logger := CreateLogger()
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err, "logger.Sync() Fatal.")
		}
	}()

	r := chi.NewRouter()
	s := NewServerImpl(logger, r)
	s.Init(*env)

	h := Openapi.HandlerFromMux(s, s.router)
	s.httpHandler = h

	log.Println("Starting app")
	_ = http.ListenAndServe(fmt.Sprint(":", *port), r)
}
