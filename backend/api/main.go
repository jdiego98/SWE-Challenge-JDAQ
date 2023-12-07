package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jdiego98/SWE-Challenge-JDAQ/api/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/api", routes.SearchRoutes())

	http.ListenAndServe(":8080", r)
}
