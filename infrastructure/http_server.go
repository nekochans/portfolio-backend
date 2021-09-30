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
	"github.com/nekochans/portfolio-backend/config"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/nekochans/portfolio-backend/usecase/memberusecase"
	"github.com/nekochans/portfolio-backend/usecase/webserviceusecase"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type HttpServer struct {
	HttpHandler http.Handler
	Router      *chi.Mux
	Db          *sql.DB
	Logger      *zap.Logger
}

func (s *HttpServer) GetMembers(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MysqlMemberRepository{Db: s.Db, Logger: s.Logger}

	u := memberusecase.UseCase{MemberRepository: repo}
	members, err := u.FetchAllFromMysql()

	if err != nil {
		if err := CreateErrorResponse(w, r, err); err != nil {
			s.Logger.Error(
				err.Error(),
				zap.Error(
					errors.Wrap(err, "failed to create error response"),
				),
			)
		}

		return
	}

	if err := CreateJsonResponse(w, r, http.StatusOK, members); err != nil {
		s.Logger.Error(
			err.Error(),
			zap.Error(
				errors.Wrap(err, "failed to create json response"),
			),
		)
	}
}

func (s *HttpServer) GetMemberById(w http.ResponseWriter, r *http.Request, id int) {
	repo := &repository.MysqlMemberRepository{Db: s.Db, Logger: s.Logger}
	u := memberusecase.UseCase{MemberRepository: repo}

	req := &memberusecase.MemberFetchRequest{Id: id}
	member, err := u.FetchFromMysql(*req)
	if err != nil {
		if err := CreateErrorResponse(w, r, err); err != nil {
			s.Logger.Error(
				err.Error(),
				zap.Error(
					errors.Wrap(err, "failed to create error response"),
				),
			)
		}
	}

	if err := CreateJsonResponse(w, r, http.StatusOK, member); err != nil {
		s.Logger.Error(
			err.Error(),
			zap.Error(
				errors.Wrap(err, "failed to create json response"),
			),
		)
	}
}

func (s *HttpServer) GetWebservices(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MysqlWebServiceRepository{Db: s.Db, Logger: s.Logger}

	u := webserviceusecase.UseCase{WebServiceRepository: repo}

	res, ErrFetchAll := u.FetchAllFromMysql()
	if ErrFetchAll != nil {
		ErrCreateError := CreateErrorResponse(w, r, ErrFetchAll)
		if ErrCreateError != nil {
			ErrCreateErrorResponse := errors.Wrap(ErrCreateError, "failed to create error response")
			s.Logger.Error(ErrCreateErrorResponse.Error(), zap.Error(ErrCreateErrorResponse))
		}
		return
	}

	ErrCreateJson := CreateJsonResponse(w, r, http.StatusOK, res)
	if ErrCreateJson != nil {
		ErrCreateJsonResponse := errors.Wrap(ErrCreateJson, "failed to create json response")
		s.Logger.Error(ErrCreateJsonResponse.Error(), zap.Error(ErrCreateJsonResponse))
	}
}

func (s *HttpServer) middleware() {
	const timeoutSecond = 60

	s.Router.Use(middleware.RequestID)
	s.Router.Use(Logger(s.Logger))
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(time.Second * timeoutSecond))
}

func (s *HttpServer) Init(env string) {
	s.middleware()
}

func NewHttpServer(l *zap.Logger, r *chi.Mux) *HttpServer {
	return &HttpServer{
		Router: r,
		Logger: l,
	}
}

func NewHttpServerWithMysql(d *sql.DB, l *zap.Logger, r *chi.Mux) *HttpServer {
	return &HttpServer{
		Router: r,
		Db:     d,
		Logger: l,
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
		if err := logger.Sync(); err != nil {
			logger.Error(
				err.Error(),
				zap.Error(
					errors.Wrap(err, "failed to logger sync"),
				),
			)
		}
	}()

	db, err := sql.Open("mysql", config.GetDsn())
	if err != nil {
		logger.Error(
			err.Error(),
			zap.Error(
				errors.Wrap(err, "unable to connect to mysql server"),
			),
		)
	}

	router := chi.NewRouter()
	server := NewHttpServerWithMysql(db, logger, router)
	server.Init(*env)

	server.HttpHandler = Openapi.HandlerFromMux(server, server.Router)

	log.Println("Starting app")
	if err := http.ListenAndServe(fmt.Sprint(":", *port), server.Router); err != nil {
		logger.Error(
			err.Error(),
			zap.Error(
				errors.Wrap(err, "unable to http listen and serve"),
			),
		)
	}
}
