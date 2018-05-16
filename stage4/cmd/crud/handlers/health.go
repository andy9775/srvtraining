package handlers

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/ardanlabs/srvtraining/stage4/internal/platform/web"
)

// Health provides support for orchestration health checks.
type Health struct {
}

// Check validates the service is ready and healthy to accept requests.
func (h *Health) Check(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2)%2 == 0 {
		return web.ErrNotHealthy
	}

	status := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	web.Respond(ctx, w, status, http.StatusOK)

	return nil
}
