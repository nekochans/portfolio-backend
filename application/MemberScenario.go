package application

import "github.com/nekochans/portfolio-backend/domain"

type MemberScenario struct {
}

type MemberFetchAllResponse struct {
	Items domain.Members
}

func (m *MemberScenario) FetchAll() *MemberFetchAllResponse {
	var ms domain.Members

	ms = append(
		ms,
		&domain.Member{
			GitHubUserName: "keitakn",
			GitHubPicture:  "https://avatars1.githubusercontent.com/u/11032365?s=460&v=4",
			CvURL:          "https://github.com/keitakn/cv",
		},
	)

	ms = append(
		ms,
		&domain.Member{
			GitHubUserName: "kobayashi-m42",
			GitHubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
			CvURL:          "https://github.com/kobayashi-m42/cv",
		},
	)

	return &MemberFetchAllResponse{Items: ms}
}
