package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nekochans/portfolio-backend/application"
	"github.com/nekochans/portfolio-backend/config"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
	// TODO DBオブジェクトの生成場所は別の場所を検討する
	db, err := sql.Open("mysql", config.GetDsn())
	log.Println(db)
	log.Println(err)
	repo := &repository.MySQLMemberRepository{DB: db}

	ms := application.MemberScenario{MemberRepository: repo}
	ml := ms.FetchAllFromMySQL()
	CreateJsonResponse(w, http.StatusOK, ml)
}
