package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

func CreateJsonResponse(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
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
	_, err = w.Write(res)
	if err != nil {
		log.Fatal(err, "http.ResponseWriter() Fatal.")
	}
}

// respondError レスポンスとして返すエラーを生成する
func CreateErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger := CreateLogger()
	logger.Error(err.Error(), zap.String("RequestId", middleware.GetReqID(r.Context())))

	hc := &HttpErrorCreator{}
	he := hc.CreateFromMsg(err.Error())
	CreateJsonResponse(w, r, he.Code, he)
}
