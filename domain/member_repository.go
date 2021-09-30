package domain

import (
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/pkg/errors"
)

var (
	ErrMemberNotFound             = errors.New("MemberRepository member not found")
	ErrMemberRepositoryUnexpected = errors.New("MemberRepository unexpected error")
)

type MemberRepository interface {
	Find(id int) (*Openapi.Member, error)
	FindAll() (Members, error)
}
