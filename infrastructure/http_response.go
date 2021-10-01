package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var ErrResponsePayloadJsonEncode = errors.New("failed to json encode of the response payload")
var ErrWriteResponse = errors.New("failed to write response payload")

func CreateJsonResponse(w http.ResponseWriter, r *http.Request, status Openapi.ErrorCode, payload interface{}) error {
	res, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(ErrResponsePayloadJsonEncode, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Request-Id", middleware.GetReqID(r.Context()))
	w.WriteHeader(int(status))
	if _, err := w.Write(res); err != nil {
		return errors.Wrap(ErrWriteResponse, err.Error())
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
