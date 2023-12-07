package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jdiego98/SWE-Challenge-JDAQ/api/handlers"
)

func SearchRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/search", handlers.SearchHandler)
	return r
}
