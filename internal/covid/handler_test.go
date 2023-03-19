package covid_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/tirathawat/covid/internal/covid"
	"github.com/tirathawat/covid/pkg/errs"
	"github.com/tirathawat/covid/pkg/ptr"
	"github.com/tirathawat/covid/pkg/testutils"
)

type mockFetcher struct {
	data *covid.Data
	err  error
}

func (f *mockFetcher) Fetch() (*covid.Data, error) {
	return f.data, f.err
}

func TestHandler_Summary(t *testing.T) {
	t.Run("Should return summary when fetch data successfully", func(t *testing.T) {
		fetcher := &mockFetcher{
			data: &covid.Data{
				Records: []covid.Record{
					{ProvinceID: ptr.Addr(1), Province: ptr.Addr("Bangkok"), Age: ptr.Addr(20)},
					{ProvinceID: ptr.Addr(2), Province: ptr.Addr("Samut Prakan"), Age: ptr.Addr(30)},
				},
			},
		}
		expected := covid.SummaryResponse{
			Province: covid.ProvinceCount{
				"Bangkok":      1,
				"Samut Prakan": 1,
			},
			AgeGroup: covid.AgeGroupCount{
				"0-30": 2,
			},
		}

		rec := testutils.HTTPRequest(covid.NewHandler(fetcher).Summary, nil)
		got := testutils.Body[covid.SummaryResponse](rec.Body.Bytes())

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %v, got %v", http.StatusOK, rec.Code)
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("Should return error when fetch data failed", func(t *testing.T) {
		fetcher := &mockFetcher{
			err: covid.ErrFetchFailed,
		}

		rec := testutils.HTTPRequest(covid.NewHandler(fetcher).Summary, nil)
		got := testutils.Body[errs.Error](rec.Body.Bytes())

		if rec.Code != http.StatusServiceUnavailable {
			t.Errorf("expected status code %v, got %v", http.StatusServiceUnavailable, rec.Code)
		}
		if got.Msg != covid.ErrFetchFailed.Msg {
			t.Errorf("expected error message %v, got %v", covid.ErrFetchFailed.Msg, got.Msg)
		}
	})
}
