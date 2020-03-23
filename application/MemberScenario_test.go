package application

import (
	"database/sql"
	"github.com/nekochans/portfolio-backend/config"
	"github.com/nekochans/portfolio-backend/domain"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	"reflect"
	"testing"
)

func createTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", config.GetTestDsn())

	if err != nil {
		t.Fatal("DB Connect Error", err)
	}

	return db
}

func fixtureTestFetchAllFromMySQLSucceed(db *sql.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS=0")
	db.Exec("TRUNCATE members")
	db.Exec("TRUNCATE members_github_users")
	db.Exec("INSERT INTO members (id) VALUE (10)")
	db.Exec("INSERT INTO members_github_users (id, member_id, github_id, avatar_url, cv_repo_name) VALUE (1, 10, 'keita', 'https://aaa.png', 'cv')")
	db.Exec("SET FOREIGN_KEY_CHECKS=1")
}

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

func TestFetchAllFromMySQLSucceed(t *testing.T) {
	db := createTestDB(t)
	fixtureTestFetchAllFromMySQLSucceed(db)

	repo := &repository.MySQLMemberRepository{DB: db}
	ms := &MemberScenario{MemberRepository: repo}
	res := ms.FetchAllFromMySQL()

	var expected domain.Members

	expected = append(
		expected,
		&domain.Member{
			ID:             10,
			GitHubUserName: "keita",
			GitHubPicture:  "https://aaa.png",
			CvURL:          "https://github.com/keita/cv",
		},
	)

	for i, member := range res.Items {
		if reflect.DeepEqual(member, expected[i]) == false {
			t.Error("\nActually: ", member, "\nExpected: ", expected[i])
		}
	}
}
