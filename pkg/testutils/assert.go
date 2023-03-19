package testutils

import (
	"bytes"
	"encoding/json"
	"testing"
)

// AssertEqualJSON compares two JSON strings.
// It will minify both JSON strings before comparing.
// It will fail the test if the two JSON strings are not equal.
func AssertEqualJSON(t *testing.T, expected string, actual interface{}) {
	t.Helper()

	actualJSON, err := json.Marshal(actual)
	if err != nil {
		t.Errorf("failed to marshal actual to json: %v", err)
	}

	expectedJSON := new(bytes.Buffer)
	if err = json.Compact(expectedJSON, []byte(expected)); err != nil {
		t.Errorf("failed to minify actual json: %v", err)
	}

	if expectedJSON.String() != string(actualJSON) {
		t.Errorf("expected %s, got %s", expectedJSON.String(), string(actualJSON))
	}
}
