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

func TestWebServiceFetchAllFromMemorySucceed(t *testing.T) {
	var expected domain.WebServices

	expected = append(
		expected,
		&domain.WebService{
			ID:          1,
			URL:         "https://www.mindexer.net",
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

func fixtureTestWebServiceFetchAllFromMySQLSucceed(t *testing.T, db *sql.DB) {
	testDataDir, err := filepath.Abs("../test/data/webservicescenario/fetchallfrommysql/succeed")
	if err != nil {
		t.Fatal("fixtureTestWebServiceFetchAllFromMySQLSucceed Error", err)
	}

	seeder := &test.Seeder{DB: db, DirPath: testDataDir}
	err = seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestWebServiceFetchAllFromMySQLSucceed Error", err)
	}

	err = seeder.Execute()
	if err != nil {
		t.Fatal("fixtureTestWebServiceFetchAllFromMySQLSucceed Error", err)
	}
}

func fixtureTestWebServiceFetchAllFromMySQLFailureWebServicesNotFound(t *testing.T, db *sql.DB) {
	seeder := &test.Seeder{DB: db}
	err := seeder.TruncateAllTable()
	if err != nil {
		t.Fatal("fixtureTestWebServiceFetchAllFromMySQLFailureWebServicesNotFound Error", err)
	}
}

func TestWebServiceFetchAllFromMySQLSucceed(t *testing.T) {
	dbCreator := &test.DBCreator{}
	db := dbCreator.Create(t)
	fixtureTestWebServiceFetchAllFromMySQLSucceed(t, db)

	repo := &repository.MySQLWebServiceRepository{DB: db}
	ws := &WebServiceScenario{WebServiceRepository: repo}
	res, err := ws.FetchAllFromMySQL()

	var expected domain.WebServices

	expected = append(
		expected,
		&domain.WebService{
			ID:          10,
			URL:         "https://stg-www.nekochans.net",
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

func TestWebServiceFetchAllFromMySQLFailureWebServicesNotFound(t *testing.T) {
	dbCreator := &test.DBCreator{}
	db := dbCreator.Create(t)
	fixtureTestWebServiceFetchAllFromMySQLFailureWebServicesNotFound(t, db)

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
