package application

import (
	"github.com/nekochans/portfolio-backend/domain"
	"reflect"
	"testing"
)

func TestFetchAllFromMemorySucceed(t *testing.T) {
	var expected domain.Members

	expected = append(
		expected,
		&domain.Member{
			ID:             1,
			GitHubUserName: "keitakn",
			GitHubPicture:  "https://avatars1.githubusercontent.com/u/11032365?s=460&v=4",
			CvURL:          "https://github.com/keitakn/cv",
		},
	)

	expected = append(
		expected,
		&domain.Member{
			ID:             2,
			GitHubUserName: "kobayashi-m42",
			GitHubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
			CvURL:          "https://github.com/kobayashi-m42/cv",
		},
	)

	ms := &MemberScenario{}
	res := ms.FetchAll()

	for i, member := range res.Items {
		if reflect.DeepEqual(member, expected[i]) == false {
			t.Error("\nActually: ", member, "\nExpected: ", expected[i])
		}
	}
}
