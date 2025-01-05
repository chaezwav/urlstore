package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"urlstore/cmd/api/resource/database"
)

func New(d *database.Postgres) *chi.Mux {
	r := chi.NewRouter()

	linkApi := Api{d}
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte(`{ "code": 404, "error": "Not found" }`))
	})

	r.Get("/{link}", linkApi.Redirect)

	r.Route("/api/link", func(r chi.Router) {
		r.Post("/", linkApi.Create) // DELETE /articles/123
	})

	return r
}
