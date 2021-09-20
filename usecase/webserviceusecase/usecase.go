package webserviceusecase

import (
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
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
		return nil, err
	}

	return &WebServiceFetchAllResponse{Items: res}, nil
}
