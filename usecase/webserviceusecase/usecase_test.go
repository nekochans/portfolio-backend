package webserviceusecase

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

func TestFetchAllHandler(t *testing.T) {
	t.Run("Success Fetch All WebServices", func(t *testing.T) {
		u := &UseCase{}
		res := u.FetchAll()

		var expected domain.WebServices

		expected = append(
			expected,
			&Openapi.WebService{
				Id:          1,
				Url:         "https://www.mindexer.net",
				Description: "This service makes Qiita stock convenient.",
			},
		)

		for i, webService := range res.Items {
			if reflect.DeepEqual(webService, expected[i]) == false {
				t.Error("\nActually: ", webService, "\nExpected: ", expected[i])
			}
		}
	})
}

func TestFetchAllFromMysqlHandler(t *testing.T) {
	t.Run("Success Fetch All WebServices", func(t *testing.T) {
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

		repo := &repository.MysqlWebServiceRepository{Db: db}
		u := &UseCase{WebServiceRepository: repo}
		res, err := u.FetchAllFromMysql()

		var expected domain.WebServices

		expected = append(
			expected,
			&Openapi.WebService{
				Id:          10,
				Url:         "https://stg-www.nekochans.net",
				Description: "This service makes Qiita stock convenient.",
			},
		)

		if err != nil {
			t.Error("\nActually: ", err, "\nExpected: ", expected)
		}

		for i, webService := range res.Items {
			if reflect.DeepEqual(webService, expected[i]) == false {
				t.Error("\nActually: ", webService, "\nExpected: ", expected[i])
			}
		}
	})

	t.Run("Error WebServices Not Found", func(t *testing.T) {
		seeder := &test.Seeder{Db: db}

		t.Cleanup(func() { _ = seeder.TruncateAllTable() })

		repo := &repository.MysqlWebServiceRepository{Db: db}
		u := &UseCase{WebServiceRepository: repo}
		res, err := u.FetchAllFromMysql()
		expected := domain.ErrWebServiceNotFound

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
