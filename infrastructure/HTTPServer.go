package infrastructure

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"time"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	return &Server{
		router: chi.NewRouter(),
	}
}

// Init 実行時にしたいこと
func (s *Server) Init(env string) {
	// 何かする
	log.Printf("env: %s", env)
}

// Middleware ミドルウェア
func (s *Server) Middleware() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(time.Second * 60))
}

// Router ルーティング設定
func (s *Server) Router() {
	h := NewHandler()

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		type json struct {
			Message string `json:"message"`
		}
		res := json{Message: "I like cat. 🐱🐱"}
		CreateJsonResponse(w, http.StatusOK, res)
	})
	s.router.Route("/members", func(members chi.Router) {
		members.Get("/{id}", h.ShowMember)
		members.Get("/", h.MemberList)
	})
}

func StartHTTPServer() {
	var (
		port = flag.String("port", "8888", "addr to bind")
		env  = flag.String("env", "develop", "実行環境 (production, staging, develop)")
	)
	flag.Parse()
	s := NewServer()
	s.Init(*env)
	s.Middleware()
	s.Router()
	log.Println("Starting app")
	http.ListenAndServe(fmt.Sprint(":", *port), s.router)
}
