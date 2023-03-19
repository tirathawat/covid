package covid_test

import (
	"testing"

	"github.com/tirathawat/covid/internal/covid"
	"github.com/tirathawat/covid/pkg/ptr"
)

func Test_ClassifyAgeGroup(t *testing.T) {
	t.Run("Should return unspecified when age is nil", func(t *testing.T) {
		actual := covid.ClassifyAgeGroup(nil)

		if actual != covid.AgeGroupUnspecified {
			t.Errorf("expected %v, got %v", covid.AgeGroupUnspecified, actual)
		}
	})

	t.Run("Should return unspecified when age is zero", func(t *testing.T) {
		age := 0
		actual := covid.ClassifyAgeGroup(&age)

		if actual != covid.AgeGroupUnspecified {
			t.Errorf("expected %v, got %v", covid.AgeGroupUnspecified, actual)
		}
	})

	t.Run("Should return unspecified when age is negative", func(t *testing.T) {
		age := -1
		actual := covid.ClassifyAgeGroup(&age)

		if actual != covid.AgeGroupUnspecified {
			t.Errorf("expected %v, got %v", covid.AgeGroupUnspecified, actual)
		}
	})

	t.Run("Should return young adults when age is between 1 and 30", func(t *testing.T) {
		ages := []int{1, 15, 30}

		for _, age := range ages {
			actual := covid.ClassifyAgeGroup(&age)

			if actual != covid.AgeGroupYoungAdults {
				t.Errorf("expected %v, got %v", covid.AgeGroupYoungAdults, actual)
			}
		}
	})

	t.Run("Should return middle age when age is between 31 and 60", func(t *testing.T) {
		ages := []int{31, 40, 60}

		for _, age := range ages {
			actual := covid.ClassifyAgeGroup(&age)

			if actual != covid.AgeGroupMiddleAged {
				t.Errorf("expected %v, got %v", covid.AgeGroupMiddleAged, actual)
			}
		}
	})

	t.Run("Should return elderly when age is between 61 and 100", func(t *testing.T) {
		ages := []int{61, 80, 100}

		for _, age := range ages {
			actual := covid.ClassifyAgeGroup(&age)

			if actual != covid.AgeGroupElderly {
				t.Errorf("expected %v, got %v", covid.AgeGroupElderly, actual)
			}
		}
	})
}

func Test_CountCasesByProvince(t *testing.T) {
	testCases := []struct {
		name     string
		input    []covid.Province
		expected covid.ProvinceCount
	}{
		{
			name:     "Should return empty map when provinces is empty",
			input:    []covid.Province{},
			expected: covid.ProvinceCount{},
		},
		{
			name: "Should return map with one province when provinces has one province",
			input: []covid.Province{
				{
					ID:   1,
					Name: "Bangkok",
				},
			},
			expected: covid.ProvinceCount{
				"Bangkok": 1,
			},
		},
		{
			name: "Should return map with two provinces when provinces has two provinces",
			input: []covid.Province{
				{
					ID:   1,
					Name: "Bangkok",
				},
				{
					ID:   2,
					Name: "Nonthaburi",
				},
				{
					ID:   1,
					Name: "Bangkok",
				},
			},
			expected: covid.ProvinceCount{
				"Bangkok":    2,
				"Nonthaburi": 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := covid.CountCasesByProvince(tc.input)

			if len(actual) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}

			for k, v := range tc.expected {
				if actual[k] != v {
					t.Errorf("expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}

func Test_CountCasesByAgeGroup(t *testing.T) {
	testCases := []struct {
		name     string
		input    []*int
		expected covid.AgeGroupCount
	}{
		{
			name:     "Should return empty map when ages is empty",
			input:    []*int{},
			expected: covid.AgeGroupCount{},
		},
		{
			name:  "Should return map with one age group when ages has one age",
			input: []*int{ptr.Addr(1), ptr.Addr(2), ptr.Addr(3)},
			expected: covid.AgeGroupCount{
				covid.AgeGroupYoungAdults: 3,
			},
		},
		{
			name:  "Should return map with two age groups when ages has two ages",
			input: []*int{ptr.Addr(1), ptr.Addr(2), ptr.Addr(3), ptr.Addr(31), ptr.Addr(32), ptr.Addr(33)},
			expected: covid.AgeGroupCount{
				covid.AgeGroupYoungAdults: 3,
				covid.AgeGroupMiddleAged:  3,
			},
		},
		{
			name:  "Should return map with three age groups when ages has three ages",
			input: []*int{ptr.Addr(1), ptr.Addr(2), ptr.Addr(3), ptr.Addr(31), ptr.Addr(32), ptr.Addr(33), ptr.Addr(61), ptr.Addr(62), ptr.Addr(63)},
			expected: covid.AgeGroupCount{
				covid.AgeGroupYoungAdults: 3,
				covid.AgeGroupMiddleAged:  3,
				covid.AgeGroupElderly:     3,
			},
		},
		{
			name:  "Should return map with three age groups when ages has three ages and unspecified age",
			input: []*int{ptr.Addr(1), ptr.Addr(2), ptr.Addr(3), ptr.Addr(31), ptr.Addr(32), ptr.Addr(33), ptr.Addr(61), ptr.Addr(62), ptr.Addr(63), nil},
			expected: covid.AgeGroupCount{
				covid.AgeGroupYoungAdults: 3,
				covid.AgeGroupMiddleAged:  3,
				covid.AgeGroupElderly:     3,
				covid.AgeGroupUnspecified: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := covid.CountCasesByAgeGroup(tc.input)

			if len(actual) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}

			for k, v := range tc.expected {
				if actual[k] != v {
					t.Errorf("expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}
