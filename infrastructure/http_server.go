package infrastructure

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nekochans/portfolio-backend/application"
	"github.com/nekochans/portfolio-backend/config"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type HttpServer struct {
	HttpHandler http.Handler
	Router      *chi.Mux
	Db          *sql.DB
	Logger      *zap.Logger
}

func (hs *HttpServer) GetMembers(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MysqlMemberRepository{Db: hs.Db}

	ms := application.MemberScenario{MemberRepository: repo}
	ml, err := ms.FetchAllFromMysql()

	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, ml)
}

func (hs *HttpServer) GetMemberById(w http.ResponseWriter, r *http.Request, id int) {
	repo := &repository.MysqlMemberRepository{Db: hs.Db}
	ms := application.MemberScenario{MemberRepository: repo}

	req := &application.MemberFetchRequest{Id: id}
	me, err := ms.FetchFromMysql(*req)
	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, me)
}

func (hs *HttpServer) GetWebservices(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MysqlWebServiceRepository{Db: hs.Db}

	ws := &application.WebServiceScenario{WebServiceRepository: repo}

	res, err := ws.FetchAllFromMysql()
	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, res)
}

func (hs *HttpServer) middleware() {
	const timeoutSecond = 60

	hs.Router.Use(middleware.RequestID)
	hs.Router.Use(Logger(hs.Logger))
	hs.Router.Use(middleware.Recoverer)
	hs.Router.Use(middleware.Timeout(time.Second * timeoutSecond))
}

func (hs *HttpServer) Init(env string) {
	hs.middleware()
}

func NewHttpServer(logger *zap.Logger, router *chi.Mux) *HttpServer {
	return &HttpServer{
		Router: router,
		Logger: logger,
	}
}

func NewHttpServerWithMysql(db *sql.DB, logger *zap.Logger, router *chi.Mux) *HttpServer {
	return &HttpServer{
		Router: router,
		Db:     db,
		Logger: logger,
	}
}

func StartHttpServer() {
	var (
		port = flag.String("port", "8888", "addr to bind")
		env  = flag.String("env", "develop", "実行環境 (production, staging, develop)")
	)
	flag.Parse()

	logger := CreateLogger()
	defer func() {
		ErrSync := logger.Sync()
		if ErrSync != nil {
			ErrLoggerSync := xerrors.Errorf("Failed to logger Sync: %w", ErrSync)
			logger.Error(ErrSync.Error(), zap.Error(ErrLoggerSync))
		}
	}()

	db, ErrConnectDb := sql.Open("mysql", config.GetDsn())
	if ErrConnectDb != nil {
		ErrMysqlConnect := xerrors.Errorf("Unable to connect to MySQL server: %w", ErrConnectDb)
		logger.Error(ErrConnectDb.Error(), zap.Error(ErrMysqlConnect))
	}

	r := chi.NewRouter()
	s := NewHttpServerWithMysql(db, logger, r)
	s.Init(*env)

	s.HttpHandler = Openapi.HandlerFromMux(s, s.Router)

	log.Println("Starting app")
	_ = http.ListenAndServe(fmt.Sprint(":", *port), s.Router)
}
