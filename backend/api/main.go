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

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Configura los encabezados CORS adecuados
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Reemplaza con tu URL de frontend
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// Llama al siguiente middleware o enrutador
			next.ServeHTTP(w, r)
		})
	})

	r.Mount("/api", routes.SearchRoutes())

	http.ListenAndServe(":8080", r)
}
