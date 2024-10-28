package routes

import (
	"io"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the root page")
}