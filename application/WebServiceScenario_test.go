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

func fixtureTestWebServiceScenarioFetchAllFromMySQLSucceed(t *testing.T, db *sql.DB) {
	testDataDir, err := filepath.Abs("../test/data/webservicescenario/fetchallfrommysql/succeed")
	if err != nil {
		t.Fatal("fixtureTestWebServiceScenarioFetchAllFromMySQLSucceed Error", err)
	}

	seeder := &test.Seeder{DB: db, DirPath: testDataDir}
	err = seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestWebServiceScenarioFetchAllFromMySQLSucceed Error", err)
	}

	err = seeder.Execute()
	if err != nil {
		t.Fatal("fixtureTestWebServiceScenarioFetchAllFromMySQLSucceed Error", err)
	}
}

func fixtureTestWebServiceScenarioFetchAllFromMySQLFailureWebServicesNotFound(t *testing.T, db *sql.DB) {
	seeder := &test.Seeder{DB: db}
	err := seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestWebServiceScenarioFetchAllFromMySQLFailureWebServicesNotFound Error", err)
	}
}

func TestWebServiceScenarioFetchAllFromMySQLSucceed(t *testing.T) {
	dbCreator := &test.DBCreator{}
	db := dbCreator.Create(t)
	fixtureTestWebServiceScenarioFetchAllFromMySQLSucceed(t, db)

	repo := &repository.MySQLWebServiceRepository{DB: db}
	ws := &WebServiceScenario{WebServiceRepository: repo}
	res, err := ws.FetchAllFromMySQL()

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

func TestWebServiceScenarioFetchAllFromMySQLFailureWebServicesNotFound(t *testing.T) {
	dbCreator := &test.DBCreator{}
	db := dbCreator.Create(t)
	fixtureTestWebServiceScenarioFetchAllFromMySQLFailureWebServicesNotFound(t, db)

	repo := &repository.MySQLWebServiceRepository{DB: db}
	ws := &WebServiceScenario{WebServiceRepository: repo}
	res, err := ws.FetchAllFromMySQL()
	expected := "MySQLWebServiceRepository.FindAll: WebServices Not Found"

	if res != nil {
		t.Error("\nActually: ", res, "\nExpected: ", expected)
	}

	if err != nil {
		if err.Error() != expected {
			t.Error("\nActually: ", err.Error(), "\nExpected: ", expected)
		}
	}
}
