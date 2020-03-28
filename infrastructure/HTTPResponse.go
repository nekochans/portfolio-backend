package infrastructure

import (
	"encoding/json"
	"fmt"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	w.Write([]byte(res))
}

// respondError レスポンスとして返すエラーを生成する
func CreateErrorResponse(w http.ResponseWriter, code int, err error) {
	fmt.Printf("%+v\n", err)
	if e, ok := err.(*HTTPError); ok {
		CreateJsonResponse(w, e.Code, e)
	} else if err != nil {
		he := HTTPError{
			Code:    code,
			Message: err.Error(),
		}
		CreateJsonResponse(w, code, he)
	}
}

// HTTPError エラー用
type HTTPError struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
}
