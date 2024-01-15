package handlers

import (
	"net/http"

	"github.com/yramanovich/runestones/log"
)

type recoveryHandler struct {
	next http.Handler
}

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx := r.Context()
			log.FromContext(ctx).ErrorContext(ctx, "recover unexpected panic", "err", err)
		}
	}()
	h.next.ServeHTTP(w, r)
}
