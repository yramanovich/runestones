package handlers

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/yramanovich/runestones/log"
)

type correlationIdHandler struct {
	next http.Handler
}

func (h correlationIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	correlationId := uuid.NewString()
	ctx := log.WithCorrelationId(r.Context(), correlationId)
	r = r.WithContext(ctx)
	h.next.ServeHTTP(w, r)
}
