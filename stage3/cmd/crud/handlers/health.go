package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// Health provides support for orchestration health checks.
type Health struct {
}

// Check validates the service is ready and healthy to accept requests.
func (h *Health) Check(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	jsonData, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		log.Printf("Respond %v Marshalling JSON response\n", err)

		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response and context.
	w.WriteHeader(http.StatusOK)

	// Send the result back to the client.
	w.Write(jsonData)

	return nil
}
