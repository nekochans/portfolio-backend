package infrastructure

import (
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/nekochans/portfolio-backend/usecase/memberusecase"
	"github.com/nekochans/portfolio-backend/usecase/webserviceusecase"
	"github.com/pkg/errors"
)

type HttpErrorCreator struct{}

func (c *HttpErrorCreator) CreateFromError(err error) Openapi.Error {
	var code Openapi.ErrorCode
	var message string

	switch errors.Cause(err) {
	case memberusecase.ErrNotFound:
		code = Openapi.ErrorCodeN404
		message = "Member Not Found"
	case webserviceusecase.ErrNotFound:
		code = Openapi.ErrorCodeN404
		message = "WebService Not Found"
	default:
		code = Openapi.ErrorCodeN500
		message = "Internal Server Error"
	}

	resErr := Openapi.Error{
		Code:    code,
		Message: message,
	}

	return resErr
}
