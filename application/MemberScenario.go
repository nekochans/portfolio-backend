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

func (m *MemberScenario) FetchFromMySQL(req MemberFetchRequest) (*Openapi.Member, error) {
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
	var ms domain.Members

	ms = append(
		ms,
		&Openapi.Member{
			Id:             1,
			GithubUserName: "keitakn",
			GithubPicture:  "https://avatars1.githubusercontent.com/u/11032365?s=460&v=4",
			CvUrl:          "https://github.com/keitakn/cv",
		},
	)

	ms = append(
		ms,
		&Openapi.Member{
			Id:             2,
			GithubUserName: "kobayashi-m42",
			GithubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
			CvUrl:          "https://github.com/kobayashi-m42/cv",
		},
	)

	return &MemberFetchAllResponse{Items: ms}
}

func (m *MemberScenario) FetchAllFromMySQL() (*MemberFetchAllResponse, error) {
	res, err := m.MemberRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return &MemberFetchAllResponse{Items: res}, nil
}
