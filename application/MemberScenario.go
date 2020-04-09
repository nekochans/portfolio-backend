package application

import (
	"github.com/nekochans/portfolio-backend/domain"
)

type MemberScenario struct {
	MemberRepository domain.MemberRepository
}

type MemberFetchRequest struct {
	MemberID int
}

func (m *MemberScenario) FetchFromMySQL(req MemberFetchRequest) (*domain.Member, error) {
	res, err := m.MemberRepository.Find(req.MemberID)
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
		&domain.Member{
			ID:             1,
			GitHubUserName: "keitakn",
			GitHubPicture:  "https://avatars1.githubusercontent.com/u/11032365?s=460&v=4",
			CvURL:          "https://github.com/keitakn/cv",
		},
	)

	ms = append(
		ms,
		&domain.Member{
			ID:             2,
			GitHubUserName: "kobayashi-m42",
			GitHubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
			CvURL:          "https://github.com/kobayashi-m42/cv",
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
