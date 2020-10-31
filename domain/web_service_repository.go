package domain

type WebServiceRepository interface {
	FindAll() (WebServices, error)
}
