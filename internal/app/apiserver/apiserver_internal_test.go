package apiserver

import (
	"testing"

	"net/http/httptest"

	"net/http"

	"github.com/stretchr/testify/assert"
)

func testAPIServerHandleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "Hello")
}
