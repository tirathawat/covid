//go:build unit

package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tirathawat/covid/cmd/api/router"
	"github.com/tirathawat/covid/config"
)

func TestRegister(t *testing.T) {
	e := router.Register(new(config.App))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	rts := e.Routes()

	if len(rts) <= 0 {
		t.Errorf("expected routes to be greater than 0, got %d", len(rts))
	}
}
