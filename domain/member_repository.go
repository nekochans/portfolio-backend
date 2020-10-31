package domain

import Openapi "github.com/nekochans/portfolio-backend/openapi"

type MemberRepository interface {
	Find(id int) (*Openapi.Member, error)
	FindAll() (Members, error)
}
