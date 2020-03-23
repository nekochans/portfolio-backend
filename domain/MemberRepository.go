package domain

type MemberRepository interface {
	FindAll() (Members, error)
}
