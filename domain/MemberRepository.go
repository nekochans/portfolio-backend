package domain

type MemberRepository interface {
	Find(memberID int) (*Member, error)
	FindAll() (Members, error)
}
