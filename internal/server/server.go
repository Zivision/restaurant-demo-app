package server

import (
	"fmt"
	"log"

	"net/http"


	"github.com/Zivision/restaurant-demo-app/internal/db"
	"github.com/Zivision/restaurant-demo-app/internal/server/routes"
)

func StartServer(PORT string) {
	db.OpenDB("test.db")
	fmt.Print(db.DB)

	// routes
	http.HandleFunc("/", routes.GetIndex)

	fmt.Printf("Server started on localhost%s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
