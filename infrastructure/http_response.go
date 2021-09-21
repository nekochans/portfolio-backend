package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

func CreateJsonResponse(w http.ResponseWriter, r *http.Request, status int, payload interface{}) error {
	res, ErrJsonEncode := json.Marshal(payload)
	if ErrJsonEncode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, ErrWriteResponse := w.Write([]byte(ErrJsonEncode.Error()))
		if ErrWriteResponse != nil {
			return ErrWriteResponse
		}
		return ErrJsonEncode
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Request-Id", middleware.GetReqID(r.Context()))
	w.WriteHeader(status)
	_, ErrWriteHeader := w.Write(res)
	if ErrWriteHeader != nil {
		return ErrWriteHeader
	}

	return nil
}

func CreateErrorResponse(w http.ResponseWriter, r *http.Request, err error) error {
	logger := CreateLogger()
	logger.Error(
		err.Error(),
		zap.String("RequestId", middleware.GetReqID(r.Context())),
		zap.Error(err),
	)

	errCreator := &HttpErrorCreator{}
	httpError := errCreator.CreateFromError(err)
	return CreateJsonResponse(w, r, httpError.Code, httpError)
}
