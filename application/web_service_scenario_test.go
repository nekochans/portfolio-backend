package application

import (
	"database/sql"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/nekochans/portfolio-backend/domain"
	"github.com/nekochans/portfolio-backend/infrastructure/repository"
	Openapi "github.com/nekochans/portfolio-backend/openapi"
	"github.com/nekochans/portfolio-backend/test"
)

func TestWebServiceScenarioFetchAllFromMemorySucceed(t *testing.T) {
	var expected domain.WebServices

	expected = append(
		expected,
		&Openapi.WebService{
			Id:          1,
			Url:         "https://www.mindexer.net",
			Description: "Qiitaのストックを便利にするサービスです。",
		},
	)

	ws := &WebServiceScenario{}
	res := ws.FetchAll()

	for i, webService := range res.Items {
		if reflect.DeepEqual(webService, expected[i]) == false {
			t.Error("\nActually: ", webService, "\nExpected: ", expected[i])
		}
	}
}

func fixtureTestWebServiceScenarioFetchAllFromMysqlSucceed(t *testing.T, db *sql.DB) {
	testDataDir, ErrFilepath := filepath.Abs("../test/data/webservicescenario/fetchallfrommysql/succeed")
	if ErrFilepath != nil {
		t.Fatal("fixtureTestWebServiceScenarioFetchAllFromMysqlSucceed Error", ErrFilepath)
	}

	seeder := &test.Seeder{Db: db, DirPath: testDataDir}
	ErrTruncate := seeder.TruncateAllTable()
	if ErrTruncate != nil {
		t.Fatal("fixtureTestWebServiceScenarioFetchAllFromMysqlSucceed Error", ErrTruncate)
	}

	ErrSeeder := seeder.Execute()
	if ErrSeeder != nil {
		t.Fatal("fixtureTestWebServiceScenarioFetchAllFromMysqlSucceed Error", ErrSeeder)
	}
}

func fixtureTestWebServiceScenarioFetchAllFromMysqlFailureWebServicesNotFound(t *testing.T, db *sql.DB) {
	seeder := &test.Seeder{Db: db}
	err := seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestWebServiceScenarioFetchAllFromMysqlFailureWebServicesNotFound Error", err)
	}
}

func TestWebServiceScenarioFetchAllFromMysqlSucceed(t *testing.T) {
	dbCreator := &test.DbCreator{}
	db := dbCreator.Create(t)
	fixtureTestWebServiceScenarioFetchAllFromMysqlSucceed(t, db)

	repo := &repository.MysqlWebServiceRepository{Db: db}
	ws := &WebServiceScenario{WebServiceRepository: repo}
	res, err := ws.FetchAllFromMysql()

	var expected domain.WebServices

	expected = append(
		expected,
		&Openapi.WebService{
			Id:          10,
			Url:         "https://stg-www.nekochans.net",
			Description: "Mindexerは、Qiitaのストックに カテゴリ機能を追加したサービスです。",
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
}

func TestWebServiceScenarioFetchAllFromMysqlFailureWebServicesNotFound(t *testing.T) {
	dbCreator := &test.DbCreator{}
	db := dbCreator.Create(t)
	fixtureTestWebServiceScenarioFetchAllFromMysqlFailureWebServicesNotFound(t, db)

	repo := &repository.MysqlWebServiceRepository{Db: db}
	ws := &WebServiceScenario{WebServiceRepository: repo}
	res, err := ws.FetchAllFromMysql()
	expected := "MysqlWebServiceRepository.FindAll: WebServices Not Found"

	if res != nil {
		t.Error("\nActually: ", res, "\nExpected: ", expected)
	}

	if err != nil {
		if err.Error() != expected {
			t.Error("\nActually: ", err.Error(), "\nExpected: ", expected)
		}
	}
}
