package application

import (
	"database/sql"
	"github.com/nekochans/portfolio-backend/domain"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
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
	dbCreator := &test.DbCreator{}
	db := dbCreator.Create(t)
	fixtureTestMemberScenarioFetchFromMySQLSucceed(t, db)

	expected := &Openapi.Member{
		Id:             1,
		GithubUserName: "keitakn",
		GithubPicture:  "https://avatars3.githubusercontent.com/u/11032365",
		CvUrl:          "https://github.com/keitakn/cv",
	}

	repo := &repository.MysqlMemberRepository{Db: db}
	ms := &MemberScenario{MemberRepository: repo}
	req := &MemberFetchRequest{Id: 1}

	res, err := ms.FetchFromMysql(*req)

	if err != nil {
		t.Error("\nActually: ", err, "\nExpected: ", expected)
	}

	if reflect.DeepEqual(res, expected) == false {
		t.Error("\nActually: ", res, "\nExpected: ", expected)
	}
}

func TestMemberScenarioFetchFromMySQLFailureMemberNotFound(t *testing.T) {
	dbCreator := &test.DbCreator{}
	db := dbCreator.Create(t)
	fixtureTestMemberScenarioFetchFromMySQLFailureMembersNotFound(t, db)

	repo := &repository.MysqlMemberRepository{Db: db}
	ms := &MemberScenario{MemberRepository: repo}
	req := &MemberFetchRequest{Id: 99}

	res, err := ms.FetchFromMysql(*req)
	expected := "MysqlMemberRepository.Find: Member Not Found"

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
		&Openapi.Member{
			Id:             1,
			GithubUserName: "keitakn",
			GithubPicture:  "https://avatars1.githubusercontent.com/u/11032365?s=460&v=4",
			CvUrl:          "https://github.com/keitakn/cv",
		},
	)

	expected = append(
		expected,
		&Openapi.Member{
			Id:             2,
			GithubUserName: "kobayashi-m42",
			GithubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
			CvUrl:          "https://github.com/kobayashi-m42/cv",
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
	dbCreator := &test.DbCreator{}
	db := dbCreator.Create(t)
	fixtureTestMemberScenarioFetchAllFromMySQLSucceed(t, db)

	repo := &repository.MysqlMemberRepository{Db: db}
	ms := &MemberScenario{MemberRepository: repo}
	res, err := ms.FetchAllFromMysql()

	var expected domain.Members

	expected = append(
		expected,
		&Openapi.Member{
			Id:             10,
			GithubUserName: "keita-kk",
			GithubPicture:  "https://aaa.png",
			CvUrl:          "https://github.com/keita/cv",
		},
		&Openapi.Member{
			Id:             20,
			GithubUserName: "moko-cat",
			GithubPicture:  "https://neko.jpeg",
			CvUrl:          "https://github.com/moko-cat/resume",
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
	dbCreator := &test.DbCreator{}
	db := dbCreator.Create(t)
	fixtureTestMemberScenarioFetchAllFromMySQLFailureMembersNotFound(t, db)

	repo := &repository.MysqlMemberRepository{Db: db}
	ms := &MemberScenario{MemberRepository: repo}
	res, err := ms.FetchAllFromMysql()
	expected := "MysqlMemberRepository.FindAll: Members Not Found"

	if res != nil {
		t.Error("\nActually: ", res, "\nExpected: ", expected)
	}

	if err != nil {
		if err.Error() != expected {
			t.Error("\nActually: ", err.Error(), "\nExpected: ", expected)
		}
	}
}
