package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nekochans/portfolio-backend/application"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	"net/http"
	"strconv"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler() *Handler {
	return &Handler{}
}

func NewHandlerWithMySQL(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) ShowMember(w http.ResponseWriter, r *http.Request) {
	type json struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res := json{ID: id, Name: fmt.Sprint("name_", id)}
	CreateJsonResponse(w, http.StatusOK, res)
}

func (h *Handler) MemberList(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MySQLMemberRepository{DB: h.DB}

	ms := application.MemberScenario{MemberRepository: repo}
	ml, err := ms.FetchAllFromMySQL()

	if err != nil {
		CreateErrorResponse(w, 400, err)
		return
	}

	CreateJsonResponse(w, http.StatusOK, ml)
}
