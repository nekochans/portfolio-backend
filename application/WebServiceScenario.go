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
	var ws domain.WebServices

	ws = append(
		ws,
		&Openapi.WebService{
			Id:          1,
			Url:         "https://www.mindexer.net",
			Description: "Qiitaのストックを便利にするサービスです。",
		},
	)

	return &WebServiceFetchAllResponse{Items: ws}
}

func (w *WebServiceScenario) FetchAllFromMySQL() (*WebServiceFetchAllResponse, error) {
	res, err := w.WebServiceRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &WebServiceFetchAllResponse{Items: res}, nil
}
