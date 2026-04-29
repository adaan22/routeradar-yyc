package api

import "net/http"

func NewRouter(corsAllowOrigin string) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)

	// Placeholder route group for future API endpoints.
	mux.HandleFunc("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	return withCORS(mux, corsAllowOrigin)
}

