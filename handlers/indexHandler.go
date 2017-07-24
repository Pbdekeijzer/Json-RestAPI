package handlers

import (
	"fmt"
	"net/http"
)

// Index is the root at "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the JSON restAPI!")
}
