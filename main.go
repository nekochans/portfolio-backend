package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server Server
type Server struct {
	router *chi.Mux
}

// New Server構造体のコンストラクタ
func New() *Server {
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
		res := json{ Message: "I like cat. 🐱🐱" }
		respondJSON(w, http.StatusOK, res)
	})
	s.router.Route("/members", func(members chi.Router) {
		members.Get("/{id}", h.ShowMember)
		members.Get("/", h.MemberList)
	})
}

// Handler ハンドラ用
type Handler struct {
}

// NewHandler コンストラクタ
func NewHandler() *Handler {
	return &Handler{}
}

// Show endpoint
func (h *Handler) ShowMember(w http.ResponseWriter, r *http.Request) {
	type json struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res := json{ID: id, Name: fmt.Sprint("name_", id)}
	respondJSON(w, http.StatusOK, res)
}

// List endpoint
func (h *Handler) MemberList(w http.ResponseWriter, r *http.Request) {
	users := []struct {
		ID   int    `json:"id"`
		User string `json:"user"`
	}{
		{1, "🐱"},
		{2, "🐶"},
		{3, "🐰"},
	}
	respondJSON(w, http.StatusOK, users)
}

// respondJSON レスポンスとして返すjsonを生成して、writerに書き込む
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError レスポンスとして返すエラーを生成する
func respondError(w http.ResponseWriter, code int, err error) {
	log.Printf("err: %v", err)
	if e, ok := err.(*HTTPError); ok {
		respondJSON(w, e.Code, e)
	} else if err != nil {
		he := HTTPError{
			Code:    code,
			Message: err.Error(),
		}
		respondJSON(w, code, he)
	}
}

// HTTPError エラー用
type HTTPError struct {
	Code    int
	Message string
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
}

func main() {
	var (
		port = flag.String("port", "8888", "addr to bind")
		env  = flag.String("env", "develop", "実行環境 (production, staging, develop)")
	)
	flag.Parse()
	s := New()
	s.Init(*env)
	s.Middleware()
	s.Router()
	log.Println("Starting app")
	http.ListenAndServe(fmt.Sprint(":", *port), s.router)
}
