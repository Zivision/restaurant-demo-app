package routes

import (
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the index page"))
}
