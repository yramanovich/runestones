package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yramanovich/runestones/runestones"
)

func SetupHandlers(m *runestones.Manager) *mux.Router {
	router := mux.NewRouter()
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
	return router
}
