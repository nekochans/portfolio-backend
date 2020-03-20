package domain

type Member struct {
	GitHubUserName string
	GitHubPicture  string
	CvURL          string
}

type Members []*Member
