package application

import (
	"github.com/nekochans/portfolio-backend/domain"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
)

type MemberScenario struct {
	MemberRepository domain.MemberRepository
}

type MemberFetchRequest struct {
	Id int
}

func (m *MemberScenario) FetchFromMysql(req MemberFetchRequest) (*Openapi.Member, error) {
	res, err := m.MemberRepository.Find(req.Id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MemberFetchAllResponse struct {
	Items domain.Members `json:"items"`
}

func (m *MemberScenario) FetchAll() *MemberFetchAllResponse {
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

func (m *MemberScenario) FetchAllFromMysql() (*MemberFetchAllResponse, error) {
	res, err := m.MemberRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &MemberFetchAllResponse{Items: res}, nil
}
