package routes

import "net/http"

func HandleRoutes() {
	http.HandleFunc("/", getRoot)
}