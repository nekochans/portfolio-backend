package webserviceusecase

import (
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/pkg/errors"
)

var (
	ErrNotFound   = errors.New("WebServiceUseCase web service not found")
	ErrUnexpected = errors.New("WebServiceUseCase unexpected error")
)

type UseCase struct {
	WebServiceRepository domain.WebServiceRepository
}

type WebServiceFetchAllResponse struct {
	Items domain.WebServices `json:"items"`
}

func (u *UseCase) FetchAll() *WebServiceFetchAllResponse {
	var item domain.WebServices

	item = append(
		item,
		&Openapi.WebService{
			Id:          1,
			Url:         "https://www.mindexer.net",
			Description: "This service makes Qiita stock convenient.",
		},
	)

	return &WebServiceFetchAllResponse{Items: item}
}

func (u *UseCase) FetchAllFromMysql() (*WebServiceFetchAllResponse, error) {
	res, err := u.WebServiceRepository.FindAll()
	if err != nil {
		switch errors.Cause(err) {
		case domain.ErrWebServiceNotFound:
			return nil, errors.Wrap(ErrNotFound, err.Error())
		default:
			return nil, errors.Wrap(ErrUnexpected, err.Error())
		}
	}

	return &WebServiceFetchAllResponse{Items: res}, nil
}
