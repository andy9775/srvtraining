package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ardanlabs/srvtraining/stage2/internal/platform/web"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	app := web.New()

	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request, p map[string]string) error {
		fmt.Fprintln(w, "test")
		return nil
	}
	app.Handle("GET", "/test", web.Handler(f))

	return app
}
