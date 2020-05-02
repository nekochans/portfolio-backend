package domain

type WebService struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type WebServices []*WebService
