package domain

type Member struct {
	GitHubUserName string `json:"githubUserName"`
	GitHubPicture  string `json:"githubPicture"`
	CvURL          string `json:"cvUrl"`
}

type Members []*Member
