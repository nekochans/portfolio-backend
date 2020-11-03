package application

import (
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
)

type WebServiceScenario struct {
	WebServiceRepository domain.WebServiceRepository
}

type WebServiceFetchAllResponse struct {
	Items domain.WebServices `json:"items"`
}

func (w *WebServiceScenario) FetchAll() *WebServiceFetchAllResponse {
	var item domain.WebServices

	item = append(
		item,
		&Openapi.WebService{
			Id:          1,
			Url:         "https://www.mindexer.net",
			Description: "Qiitaのストックを便利にするサービスです。",
		},
	)

	return &WebServiceFetchAllResponse{Items: item}
}

func (w *WebServiceScenario) FetchAllFromMysql() (*WebServiceFetchAllResponse, error) {
	res, err := w.WebServiceRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &WebServiceFetchAllResponse{Items: res}, nil
}
