package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yramanovich/runestones/log"
)

const maxBytesLimit = 2_000_000

// RunestoneService manages runestones objects in the system.
type RunestoneService interface {
	CreateRunestone(ctx context.Context, content []byte) (string, error)
	GetRunestone(ctx context.Context, id string) ([]byte, error)
}

// SetupHandlers configures router instance with the given RunestoneService implementation.
func SetupHandlers(svc RunestoneService) *mux.Router {
	router := mux.NewRouter()
	router.
		Methods(http.MethodGet).
		Path("/v1/runestones/{id}").
		Handler(middlewares(getRunestoneHandler(svc)))

	router.
		Methods(http.MethodPost).
		Path("/v1/runestones").
		Handler(middlewares(createRunestoneHandler(svc)))
	return router
}

func middlewares(h http.Handler) http.Handler {
	h = recoveryHandler{next: h}
	h = correlationIdHandler{next: h}
	h = http.MaxBytesHandler(h, maxBytesLimit)
	return h
}

func getRunestoneHandler(svc RunestoneService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok || id == "" {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		content, err := svc.GetRunestone(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func createRunestoneHandler(svc RunestoneService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		ctx := r.Context()

		log.Close(ctx, r.Body, "close request body")

		id, err := svc.CreateRunestone(ctx, body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
