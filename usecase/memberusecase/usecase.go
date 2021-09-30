package memberusecase

import (
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/pkg/errors"
)

var (
	ErrNotFound   = errors.New("member not found")
	ErrUnexpected = errors.New("member UseCase unexpected error")
)

type UseCase struct {
	MemberRepository domain.MemberRepository
}

type MemberFetchRequest struct {
	Id int
}

func (u *UseCase) FetchFromMysql(req MemberFetchRequest) (*Openapi.Member, error) {
	res, err := u.MemberRepository.Find(req.Id)
	if err != nil {
		switch errors.Cause(err) {
		case domain.ErrMemberNotFound:
			return nil, errors.Wrap(ErrNotFound, err.Error())
		default:
			return nil, errors.Wrap(ErrUnexpected, err.Error())
		}
	}

	return res, nil
}

type MemberFetchAllResponse struct {
	Items domain.Members `json:"items"`
}

func (u *UseCase) FetchAll() *MemberFetchAllResponse {
	var item domain.Members

	const keitaMemberId = 1
	const mopMemberId = 2

	item = append(
		item,
		&Openapi.Member{
			Id:             keitaMemberId,
			GithubUserName: "keitakn",
			GithubPicture:  "https://avatars1.githubusercontent.com/u/11032365?s=460&v=4",
			CvUrl:          "https://github.com/keitakn/cv",
		},
	)

	item = append(
		item,
		&Openapi.Member{
			Id:             mopMemberId,
			GithubUserName: "kobayashi-m42",
			GithubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
			CvUrl:          "https://github.com/kobayashi-m42/cv",
		},
	)

	return &MemberFetchAllResponse{Items: item}
}

func (u *UseCase) FetchAllFromMysql() (*MemberFetchAllResponse, error) {
	res, err := u.MemberRepository.FindAll()
	if err != nil {
		switch errors.Cause(err) {
		case domain.ErrMemberNotFound:
			return nil, errors.Wrap(ErrNotFound, err.Error())
		default:
			return nil, errors.Wrap(ErrUnexpected, err.Error())
		}
	}

	return &MemberFetchAllResponse{Items: res}, nil
}
