package testutils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// HTTPRequest creates a new HTTP request with the given handler and body and returns the response.
// It is used for testing the handler.
func HTTPRequest(handler gin.HandlerFunc, body io.Reader) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(rec)
	context.Request = &http.Request{Body: io.NopCloser(body)}
	handler(context)
	return rec
}

// Body parses the response body and returns the value.
// It panics if the response body is not valid JSON.
func Body[T any](b []byte) T {
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		panic(err)
	}
	return v
}
