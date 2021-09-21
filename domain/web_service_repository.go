package domain

import "github.com/pkg/errors"

var (
	ErrWebServiceNotFound             = errors.New("web service not found")
	ErrWebServiceRepositoryUnexpected = errors.New("web service repository unexpected error")
)

type WebServiceRepository interface {
	FindAll() (WebServices, error)
}
