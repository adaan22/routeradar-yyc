package api

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// HEAD requests should return headers only.
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
		return
	}

	_ = json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
}

