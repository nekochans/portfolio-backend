package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

func CreateJsonResponse(w http.ResponseWriter, r *http.Request, status int, payload interface{}) {
	res, ErrJsonEncode := json.MarshalIndent(payload, "", "    ")
	if ErrJsonEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, ErrWriteResponse := w.Write([]byte(ErrJsonEncode.Error()))
		if ErrWriteResponse != nil {
			log.Fatal(ErrWriteResponse, "http.ResponseWriter() Fatal.")
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Request-Id", middleware.GetReqID(r.Context()))
	w.WriteHeader(status)
	_, ErrWriteHeader := w.Write(res)
	if ErrWriteHeader != nil {
		log.Fatal(ErrWriteHeader, "http.ResponseWriter() Fatal.")
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
