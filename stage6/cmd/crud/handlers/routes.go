package handlers

import (
	"net/http"

	"github.com/ardanlabs/srvtraining/stage6/internal/mid"
	"github.com/ardanlabs/srvtraining/stage6/internal/platform/web"
)

// API returns a handler for a set of routes.
func API() http.Handler {
	app := web.New(mid.RequestLogger, mid.ErrorHandler)

	h := Health{}
	app.Handle("GET", "/v1/health", h.Check)

	return app
}
