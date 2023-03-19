package covid_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tirathawat/covid/internal/covid"
	"github.com/tirathawat/covid/pkg/testutils"
)

const data = `
{
	"Data": [
		{
			"ConfirmDate": "2020-01-01",
			"No": 1,
			"Age": 50,
			"Gender": "ชาย",
			"GenderEn": "Male",
			"Nation": "China",
			"NationEn": "China",
			"Province": "กรุงเทพมหานคร",
			"ProvinceId": 1,
			"District": "บางกอกใหญ่",
			"ProvinceEn": "Bangkok",
			"StatQuarantine": 0
		},
		{
			"ConfirmDate": "2020-01-02",
			"No": 2,
			"Age": 30,
			"Gender": "หญิง",
			"GenderEn": "Female",
			"Nation": "Thailand",
			"NationEn": "Thailand",
			"Province": "สมุทรปราการ",
			"ProvinceId": 2,
			"District": "เมืองสมุทรปราการ",
			"ProvinceEn": "Samut Prakan",
			"StatQuarantine": 1
		}
	]
}
`

func FakeServer(data string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data))
	}))
}

func TestClient_Fetch(t *testing.T) {
	t.Run("Should return data when fetch successfully", func(t *testing.T) {
		server := FakeServer(data, http.StatusOK)
		defer server.Close()
		fetcher := covid.NewClient(server.URL)

		got, err := fetcher.Fetch()
		if err != nil {
			t.Fatalf("Fetch() returned an error: %v", err)
		}

		testutils.AssertEqualJSON(t, data, got)
	})

	t.Run("Should return error when fetch failed with 500 status code", func(t *testing.T) {
		expected := "unexpected status code: 500, body: cannot fetch data"
		server := FakeServer("cannot fetch data", http.StatusInternalServerError)
		defer server.Close()
		fetcher := covid.NewClient(server.URL)

		_, err := fetcher.Fetch()
		if err.Error() != expected {
			t.Fatalf("want %s, got %s", expected, err.Error())
		}
	})

	t.Run("Should return error when fetch failed with invalid url", func(t *testing.T) {
		fetcher := covid.NewClient("invalid url")

		_, err := fetcher.Fetch()
		if err == nil {
			t.Fatalf("Fetch() should return an error")
		}
	})

	t.Run("Should return error when fetch invalid response body", func(t *testing.T) {
		server := FakeServer("invalid response body", http.StatusOK)
		defer server.Close()
		fetcher := covid.NewClient(server.URL)

		_, err := fetcher.Fetch()
		if err == nil {
			t.Fatalf("Fetch() should return an error")
		}
	})
}
