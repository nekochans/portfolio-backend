package infrastructure

import (
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/nekochans/portfolio-backend/usecase/memberusecase"
	"github.com/nekochans/portfolio-backend/usecase/webserviceusecase"
	"github.com/pkg/errors"
)

type HttpErrorCreator struct{}

func (c *HttpErrorCreator) CreateFromError(err error) Openapi.Error {
	const notFoundErrorCode = 404
	const internalServerErrorCode = 500

	var code int
	var message string

	switch errors.Cause(err) {
	case memberusecase.ErrNotFound:
		code = notFoundErrorCode
		message = "Member Not Found"
	case webserviceusecase.ErrNotFound:
		code = notFoundErrorCode
		message = "WebService Not Found"
	default:
		code = internalServerErrorCode
		message = "Internal Server Error"
	}

	resErr := Openapi.Error{
		Code:    code,
		Message: message,
	}

	return resErr
}
