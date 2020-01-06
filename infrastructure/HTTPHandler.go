package infrastructure

import (
	"fmt"
	"github.com/go-chi/chi"
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
	users := []struct {
		ID   int    `json:"id"`
		User string `json:"user"`
	}{
		{1, "🐱"},
		{2, "🐶"},
		{3, "🐰"},
		{4, "(=^・^=)"},
	}
	CreateJsonResponse(w, http.StatusOK, users)
}
