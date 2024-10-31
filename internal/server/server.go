package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	//"github.com/Zivision/restaurant-demo-app/internal/db"
	"github.com/Zivision/restaurant-demo-app/internal/server/routes"
)

func StartServer(PORT string) {
	r := chi.NewRouter()
	// Middleware
	r.Use(middleware.RequestID)
 	r.Use(middleware.RealIP)
  	r.Use(middleware.Logger)
  	r.Use(middleware.Recoverer)

	r.Handle("/*", http.FileServer(http.Dir("client/dist")))

	// Index page
	r.Get("/blank", routes.GetIndex)
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}
