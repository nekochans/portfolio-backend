package domain

type Member struct {
	ID             int    `json:"id"`
	GitHubUserName string `json:"githubUserName"`
	GitHubPicture  string `json:"githubPicture"`
	CvURL          string `json:"cvUrl"`
}

type Members []*Member
