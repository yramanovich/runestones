package main

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/yramanovich/runestones/manager"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Hello World!")

	router := mux.NewRouter()

	m := manager.NewManager()

	router.Methods("GET").Path("/v1/runestones/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if id, ok := vars["id"]; ok {
			content, err := m.GetRunestone(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Write(content)
			return
		}
		http.Error(w, "id is not found", 500)
	})
	router.Methods("POST").Path("/v1/runestones").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer r.Body.Close()
		id, err := m.CreateRunestone(r.Context(), data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte(id))
	})

	server := &http.Server{
		Addr:              ":8000",
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
		BaseContext: func(net.Listener) context.Context {
			return context.WithValue(context.Background(), "logger", logger)
		},
	}

	server.ListenAndServe()
}
