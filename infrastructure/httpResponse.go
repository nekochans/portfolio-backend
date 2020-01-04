package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	res, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(res))
}

// respondError レスポンスとして返すエラーを生成する
func CreateErrorResponse(w http.ResponseWriter, code int, err error) {
	log.Printf("err: %v", err)
	if e, ok := err.(*HttpError); ok {
		CreateJsonResponse(w, e.Code, e)
	} else if err != nil {
		he := HttpError{
			Code:    code,
			Message: err.Error(),
		}
		CreateJsonResponse(w, code, he)
	}
}

// HttpError エラー用
type HttpError struct {
	Code    int
	Message string
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
}
