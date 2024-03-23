package server

import (
	"net/http"
)

// hello handles the index route.
func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}
