package application

import (
	"github.com/nekochans/portfolio-backend/domain"
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
