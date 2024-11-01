package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer(PORT string) {
	r := chi.NewRouter()
	// Middleware
	r.Use(middleware.RequestID)
 	r.Use(middleware.RealIP)
  	r.Use(middleware.Logger)
  	r.Use(middleware.Recoverer)

	r.Handle("/*", http.FileServer(http.Dir("client/dist")))

	// Start and serve server
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}
