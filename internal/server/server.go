package server

import (
	"net/http"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	//"github.com/Zivision/restaurant-demo-app/internal/db"
	"github.com/Zivision/restaurant-demo-app/internal/server/routes"
)

func StartServer(PORT string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Index page
	r.Get("/", routes.GetIndex)
	http.ListenAndServe(PORT, r)
}
