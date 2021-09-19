package memberusecase

import (
	"database/sql"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/nekochans/portfolio-backend/test"
)

var db *sql.DB

func TestMain(m *testing.M) {
	dbCreator := &test.DbCreator{}
	db, _ = dbCreator.Create()

	seeder := &test.Seeder{Db: db}
	_ = seeder.TruncateAllTable()

	status := m.Run()

	os.Exit(status)
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
}
