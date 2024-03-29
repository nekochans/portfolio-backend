package memberusecase

import (
	"database/sql"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/nekochans/portfolio-backend/domain"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/nekochans/portfolio-backend/test"
	"github.com/pkg/errors"
)

var db *sql.DB

// Go1.15 から TestMain には os.Exit() のコールが不要になったのでlintのルールを無効化
//nolint:staticcheck
func TestMain(m *testing.M) {
	dbCreator := &test.DbCreator{}
	db, _ = dbCreator.Create()

	seeder := &test.Seeder{Db: db}
	_ = seeder.TruncateAllTable()

	m.Run()

	_ = seeder.TruncateAllTable()
}

func TestFetchFromMysqlHandler(t *testing.T) {
	t.Run("Success Fetch Member", func(t *testing.T) {
		testDataDir, err := filepath.Abs("./testdata/fetchfrommysql/success")
		if err != nil {
			t.Fatal("Failed Read test data", err)
		}

		seeder := &test.Seeder{Db: db, DirPath: testDataDir}

		err = seeder.Execute()
		if err != nil {
			t.Fatal("Failed seeder.Execute()", err)
		}

		t.Cleanup(func() { _ = seeder.TruncateAllTable() })

		expected := &Openapi.Member{
			Id:             1,
			GithubUserName: "keitakn",
			GithubPicture:  "https://avatars3.githubusercontent.com/u/11032365",
			CvUrl:          "https://github.com/keitakn/cv",
		}

		repo := &repository.MysqlMemberRepository{Db: db}
		u := &UseCase{MemberRepository: repo}
		req := &MemberFetchRequest{Id: 1}

		res, err := u.FetchFromMysql(*req)

		if err != nil {
			t.Error("\nActually: ", err, "\nExpected: ", expected)
		}

		if reflect.DeepEqual(res, expected) == false {
			t.Error("\nActually: ", res, "\nExpected: ", expected)
		}
	})

	t.Run("Error Member Not Found", func(t *testing.T) {
		seeder := &test.Seeder{Db: db}

		t.Cleanup(func() { _ = seeder.TruncateAllTable() })

		repo := &repository.MysqlMemberRepository{Db: db}
		u := &UseCase{MemberRepository: repo}
		req := &MemberFetchRequest{Id: 99}

		res, err := u.FetchFromMysql(*req)
		expected := ErrNotFound

		if res != nil {
			t.Error("\nActually: ", res, "\nExpected: ", expected)
		}

		resErr := errors.Cause(err)

		if err != nil {
			if resErr != expected {
				t.Error("\nActually: ", resErr, "\nExpected: ", expected)
			}
		}
	})
}

func TestFetchAllHandler(t *testing.T) {
	t.Run("Success Fetch Members", func(t *testing.T) {
		var expected domain.Members

		expected = append(
			expected,
			&Openapi.Member{
				Id:             1,
				GithubUserName: "keitakn",
				GithubPicture:  "https://avatars1.githubusercontent.com/u/11032365?s=460&v=4",
				CvUrl:          "https://github.com/keitakn/cv",
			},
			&Openapi.Member{
				Id:             2,
				GithubUserName: "kobayashi-m42",
				GithubPicture:  "https://avatars0.githubusercontent.com/u/32682645?s=460&v=4",
				CvUrl:          "https://github.com/kobayashi-m42/cv",
			},
		)

		u := &UseCase{}
		res := u.FetchAll()

		for i, member := range res.Items {
			if reflect.DeepEqual(member, expected[i]) == false {
				t.Error("\nActually: ", member, "\nExpected: ", expected[i])
			}
		}
	})
}

//nolint:funlen
func TestFetchAllFromMysqlHandler(t *testing.T) {
	t.Run("Success Fetch All Members", func(t *testing.T) {
		testDataDir, err := filepath.Abs("./testdata/fetchallfrommysql/success")
		if err != nil {
			t.Fatal("Failed Read test data", err)
		}

		seeder := &test.Seeder{Db: db, DirPath: testDataDir}

		err = seeder.Execute()
		if err != nil {
			t.Fatal("Failed seeder.Execute()", err)
		}

		t.Cleanup(func() { _ = seeder.TruncateAllTable() })

		var expected domain.Members

		expected = append(
			expected,
			&Openapi.Member{
				Id:             10,
				GithubUserName: "keita",
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

		repo := &repository.MysqlMemberRepository{Db: db}
		u := &UseCase{MemberRepository: repo}

		res, err := u.FetchAllFromMysql()

		if err != nil {
			t.Error("\nActually: ", err, "\nExpected: ", expected)
		}

		for i, member := range res.Items {
			if reflect.DeepEqual(member, expected[i]) == false {
				t.Error("\nActually: ", member, "\nExpected: ", expected[i])
			}
		}
	})

	t.Run("Error Members Not Found", func(t *testing.T) {
		seeder := &test.Seeder{Db: db}

		t.Cleanup(func() { _ = seeder.TruncateAllTable() })

		repo := &repository.MysqlMemberRepository{Db: db}
		u := &UseCase{MemberRepository: repo}

		res, err := u.FetchAllFromMysql()
		expected := ErrNotFound

		if res != nil {
			t.Error("\nActually: ", res, "\nExpected: ", expected)
		}

		resErr := errors.Cause(err)

		if err != nil {
			if resErr != expected {
				t.Error("\nActually: ", resErr, "\nExpected: ", expected)
			}
		}
	})
}
