package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/yramanovich/runestones/log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestCorrelationIdHandler_ServeHTTP(t *testing.T) {
	var correlationId string
	h := correlationIdHandler{next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId = log.CorrelationId(r.Context())
		w.WriteHeader(http.StatusOK)
	})}

	req, _ := http.NewRequest("GET", "localhost:80/", nil)
	h.ServeHTTP(httptest.NewRecorder(), req)

	assert.Regexp(t, regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"), correlationId)
}
