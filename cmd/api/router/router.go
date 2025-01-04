package router

import (
	"github.com/go-chi/chi/v5"
	"urlstore/cmd/api/resource/link"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{id}", link.Redirect)
	r.Get("/api/link/{id}", link.Fetch)

	return r
}
