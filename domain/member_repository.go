package domain

import (
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/pkg/errors"
)

var (
	ErrMemberNotFound             = errors.New("member not found")
	ErrMemberRepositoryUnexpected = errors.New("member repository unexpected error")
)

type MemberRepository interface {
	Find(id int) (*Openapi.Member, error)
	FindAll() (Members, error)
}
