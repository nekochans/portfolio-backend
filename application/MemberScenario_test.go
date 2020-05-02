package application

import (
	"database/sql"
	"github.com/nekochans/portfolio-backend/domain"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	"github.com/nekochans/portfolio-backend/test"
	"path/filepath"
	"reflect"
	"testing"
)

func fixtureTestMemberScenarioFetchFromMySQLSucceed(t *testing.T, db *sql.DB) {
	testDataDir, err := filepath.Abs("../test/data/memberscenario/fetchfrommysql/succeed")
	if err != nil {
		t.Fatal("fixtureTestMemberScenarioFetchFromMySQLSucceed Error", err)
	}

	seeder := &test.Seeder{DB: db, DirPath: testDataDir}
	err = seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestMemberScenarioFetchFromMySQLSucceed Error", err)
	}

	err = seeder.Execute()
	if err != nil {
		t.Fatal("fixtureTestMemberScenarioFetchFromMySQLSucceed Error", err)
	}
}

func fixtureTestMemberScenarioFetchFromMySQLFailureMembersNotFound(t *testing.T, db *sql.DB) {
	seeder := &test.Seeder{DB: db}
	err := seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestMemberScenarioFetchFromMySQLFailureMembersNotFound Error", err)
	}
}

func fixtureTestMemberScenarioFetchAllFromMySQLSucceed(t *testing.T, db *sql.DB) {
	testDataDir, err := filepath.Abs("../test/data/memberscenario/fetchallfrommysql/succeed")
	if err != nil {
		t.Fatal("fixtureTestMemberScenarioFetchAllFromMySQLSucceed Error", err)
	}

	seeder := &test.Seeder{DB: db, DirPath: testDataDir}
	err = seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestMemberScenarioFetchAllFromMySQLSucceed Error", err)
	}

	err = seeder.Execute()
	if err != nil {
		t.Fatal("fixtureTestMemberScenarioFetchAllFromMySQLSucceed Error", err)
	}
}

func fixtureTestMemberScenarioFetchAllFromMySQLFailureMembersNotFound(t *testing.T, db *sql.DB) {
	seeder := &test.Seeder{DB: db}
	err := seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestMemberScenarioFetchAllFromMySQLFailureMembersNotFound Error", err)
	}
}

func TestMemberScenarioFetchFromMySQLSucceed(t *testing.T) {
	dbCreator := &test.DBCreator{}
	db := dbCreator.Create(t)
	fixtureTestMemberScenarioFetchFromMySQLSucceed(t, db)

	expected := &domain.Member{
		ID:             1,
		GitHubUserName: "keitakn",
		GitHubPicture:  "https://avatars3.githubusercontent.com/u/11032365",
		CvURL:          "https://github.com/keitakn/cv",
	}

	repo := &repository.MySQLMemberRepository{DB: db}
	ms := &MemberScenario{MemberRepository: repo}
	req := &MemberFetchRequest{MemberID: 1}

	res, err := ms.FetchFromMySQL(*req)

	if err != nil {
		t.Error("\nActually: ", err, "\nExpected: ", expected)
	}

	if reflect.DeepEqual(res, expected) == false {
		t.Error("\nActually: ", res, "\nExpected: ", expected)
	}
}

func TestMemberScenarioFetchFromMySQLFailureMemberNotFound(t *testing.T) {
	dbCreator := &test.DBCreator{}
	db := dbCreator.Create(t)
	fixtureTestMemberScenarioFetchFromMySQLFailureMembersNotFound(t, db)

	repo := &repository.MySQLMemberRepository{DB: db}
	ms := &MemberScenario{MemberRepository: repo}
	req := &MemberFetchRequest{MemberID: 99}

	res, err := ms.FetchFromMySQL(*req)
	expected := "MySQLMemberRepository.Find: Member Not Found"

	if res != nil {
		t.Error("\nActually: ", res, "\nExpected: ", expected)
	}

	if err != nil {
		if err.Error() != expected {
			t.Error("\nActually: ", err.Error(), "\nExpected: ", expected)
		}
	}
}

func TestMemberScenarioFetchAllMemorySucceed(t *testing.T) {
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

func TestMemberScenarioFetchAllFromMySQLSucceed(t *testing.T) {
	dbCreator := &test.DBCreator{}
	db := dbCreator.Create(t)
	fixtureTestMemberScenarioFetchAllFromMySQLSucceed(t, db)

	repo := &repository.MySQLMemberRepository{DB: db}
	ms := &MemberScenario{MemberRepository: repo}
	res, err := ms.FetchAllFromMySQL()

	var expected domain.Members

	expected = append(
		expected,
		&domain.Member{
			ID:             10,
			GitHubUserName: "keita",
			GitHubPicture:  "https://aaa.png",
			CvURL:          "https://github.com/keita/cv",
		},
		&domain.Member{
			ID:             20,
			GitHubUserName: "moko-cat",
			GitHubPicture:  "https://neko.jpeg",
			CvURL:          "https://github.com/moko-cat/resume",
		},
	)

	if err != nil {
		t.Error("\nActually: ", err, "\nExpected: ", expected)
	}

	for i, member := range res.Items {
		if reflect.DeepEqual(member, expected[i]) == false {
			t.Error("\nActually: ", member, "\nExpected: ", expected[i])
		}
	}
}

func TestMemberScenarioFetchAllFromMySQLFailureMembersNotFound(t *testing.T) {
	dbCreator := &test.DBCreator{}
	db := dbCreator.Create(t)
	fixtureTestMemberScenarioFetchAllFromMySQLFailureMembersNotFound(t, db)

	repo := &repository.MySQLMemberRepository{DB: db}
	ms := &MemberScenario{MemberRepository: repo}
	res, err := ms.FetchAllFromMySQL()
	expected := "MySQLMemberRepository.FindAll: Members Not Found"

	if res != nil {
		t.Error("\nActually: ", res, "\nExpected: ", expected)
	}

	if err != nil {
		if err.Error() != expected {
			t.Error("\nActually: ", err.Error(), "\nExpected: ", expected)
		}
	}
}
