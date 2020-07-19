package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"log"
	"net/http"
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
	_, err = w.Write([]byte(res))
	if err != nil {
		log.Fatal(err, "http.ResponseWriter() Fatal.")
	}
}

// respondError レスポンスとして返すエラーを生成する
func CreateErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger := CreateLogger()
	logger.Error(err.Error(), zap.String("RequestId", middleware.GetReqID(r.Context())))

	hc := &HTTPErrorCreator{}
	he := hc.CreateFromMsg(err.Error())
	CreateJsonResponse(w, r, he.Code, he)
}

// HTTPError エラー用
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
}
