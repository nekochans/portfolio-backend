package infrastructure

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nekochans/portfolio-backend/application"
	"github.com/nekochans/portfolio-backend/config"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
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

func(si *ServerImpl) GetMembers(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MySQLMemberRepository{DB: si.DB}

	ms := application.MemberScenario{MemberRepository: repo}
	ml, err := ms.FetchAllFromMySQL()

	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, ml)
}

func(si *ServerImpl) GetMemberById(w http.ResponseWriter, r *http.Request) {
	si.httpHandler = Openapi.GetMemberByIdCtx(si.httpHandler)

	id := r.Context().Value("id").(int)

	repo := &repository.MySQLMemberRepository{DB: si.DB}
	ms := application.MemberScenario{MemberRepository: repo}

	req := &application.MemberFetchRequest{Id: id}
	me, err := ms.FetchFromMySQL(*req)
	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, me)
}

func(si *ServerImpl) GetWebservices(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MySQLWebServiceRepository{DB: si.DB}

	ws := &application.WebServiceScenario{WebServiceRepository: repo}

	res, err := ws.FetchAllFromMySQL()
	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, res)
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

	db, err := sql.Open("mysql", config.GetDsn())
	if err != nil {
		log.Fatal(db, "Unable to connect to MySQL server.")
	}

	r := chi.NewRouter()
	s := NewServerImplWithMySQL(db, logger, r)
	s.Init(*env)

	s.httpHandler = Openapi.HandlerFromMux(s, s.router)

	log.Println("Starting app")
	_ = http.ListenAndServe(fmt.Sprint(":", *port), s.router)
}
