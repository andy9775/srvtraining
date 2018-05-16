package handlers

import (
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	mux := httptreemux.New()

	f := func(w http.ResponseWriter, r *http.Request, p map[string]string) {
		fmt.Fprintln(w, "test")
	}
	mux.Handle("GET", "/test", f)

	return mux
}
