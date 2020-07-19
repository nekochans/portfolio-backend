package infrastructure

import (
	"database/sql"
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
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	repo := &repository.MySQLMemberRepository{DB: h.DB}
	ms := application.MemberScenario{MemberRepository: repo}

	// TODO リクエストのバリデーションはどこかで実装する必要がある
	req := &application.MemberFetchRequest{Id: id}
	me, err := ms.FetchFromMySQL(*req)
	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, me)
}

func (h *Handler) MemberList(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MySQLMemberRepository{DB: h.DB}

	ms := application.MemberScenario{MemberRepository: repo}
	ml, err := ms.FetchAllFromMySQL()

	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, ml)
}

func (h *Handler) WebServiceList(w http.ResponseWriter, r *http.Request) {
	repo := &repository.MySQLWebServiceRepository{DB: h.DB}

	ws := &application.WebServiceScenario{WebServiceRepository: repo}

	res, err := ws.FetchAllFromMySQL()
	if err != nil {
		CreateErrorResponse(w, r, err)
		return
	}

	CreateJsonResponse(w, r, http.StatusOK, res)
}
