package server

import (
	"net/http"

	api "lifeos-api/internal/server/api"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	handler := NewHandler()
	api.HandlerFromMux(handler, r)

	return r
}
