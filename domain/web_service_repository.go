package domain

import "github.com/pkg/errors"

var (
	ErrWebServiceNotFound             = errors.New("WebServiceRepository web service not found")
	ErrWebServiceRepositoryUnexpected = errors.New("WebServiceRepository unexpected error")
)

type WebServiceRepository interface {
	FindAll() (WebServices, error)
}
