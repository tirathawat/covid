package covid_test

import (
	"reflect"
	"testing"

	"github.com/tirathawat/covid/internal/covid"
	"github.com/tirathawat/covid/pkg/ptr"
)

func TestData_Provices(t *testing.T) {
	testcases := []struct {
		name     string
		records  []covid.Record
		expected []covid.Province
	}{
		{
			name: "Should return provinces from records",
			records: []covid.Record{
				{ProvinceID: ptr.Addr(1), Province: ptr.Addr("Bangkok")},
				{ProvinceID: ptr.Addr(2), Province: ptr.Addr("Samut Prakan")},
			},
			expected: []covid.Province{
				{ID: 1, Name: "Bangkok"},
				{ID: 2, Name: "Samut Prakan"},
			},
		},
		{
			name: "Should ignore records that have nil ProvinceID or Province",
			records: []covid.Record{
				{ProvinceID: ptr.Addr(1), Province: ptr.Addr("Bangkok")},
				{ProvinceID: nil, Province: ptr.Addr("Samut Prakan")},
				{ProvinceID: ptr.Addr(2), Province: nil},
			},
			expected: []covid.Province{
				{ID: 1, Name: "Bangkok"},
			},
		},
		{
			name:     "Should return empty slice when records is empty",
			records:  []covid.Record{},
			expected: []covid.Province{},
		},
		{
			name:     "Should return empty slice when records is nil",
			records:  nil,
			expected: []covid.Province{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			data := covid.Data{Records: tc.records}

			actual := data.Provinces()

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestData_Ages(t *testing.T) {
	testcases := []struct {
		name     string
		records  []covid.Record
		expected []*int
	}{
		{
			name: "Should return ages from records",
			records: []covid.Record{
				{Age: ptr.Addr(1)},
				{Age: ptr.Addr(2)},
			},
			expected: []*int{ptr.Addr(1), ptr.Addr(2)},
		},
		{
			name: "Should convert nil Age to nil",
			records: []covid.Record{
				{Age: ptr.Addr(1)},
				{Age: nil},
			},
			expected: []*int{ptr.Addr(1), nil},
		},
		{
			name:     "Should return empty slice when records is empty",
			records:  []covid.Record{},
			expected: []*int{},
		},
		{
			name:     "Should return empty slice when records is nil",
			records:  nil,
			expected: []*int{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			data := covid.Data{Records: tc.records}

			actual := data.Ages()

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
