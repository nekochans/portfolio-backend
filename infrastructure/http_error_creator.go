package infrastructure

import (
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/pkg/errors"
)

type HttpErrorCreator struct{}

func (c *HttpErrorCreator) CreateFromError(err error) Openapi.Error {
	const notFoundErrorCode = 404
	const internalServerErrorCode = 500

	var code int
	var message string

	switch errors.Cause(err) {
	case domain.ErrMemberNotFound:
		code = notFoundErrorCode
		message = "Member Not Found"
	case domain.ErrWebServiceNotFound:
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
