package application

import "github.com/nekochans/portfolio-backend/domain"

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
		&domain.WebService{
			ID:          1,
			URL:         "https://www.mindexer.net",
			Description: "Qiitaのストックを便利にするサービスです。",
		},
	)

	return &WebServiceFetchAllResponse{Items: ws}
}
