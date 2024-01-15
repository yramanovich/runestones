package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoveryHandler_ServeHTTP(t *testing.T) {
	h := recoveryHandler{next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("unexpected panic")
	})}
	assert.NotPanics(t, func() {
		req, _ := http.NewRequest("GET", "localhost:80/", nil)
		h.ServeHTTP(httptest.NewRecorder(), req)
	})
}
