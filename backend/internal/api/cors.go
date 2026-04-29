package api

import (
	"net/http"
	"strings"
)

func withCORS(next http.Handler, allowOrigin string) http.Handler {
	// Minimal CORS middleware for Expo/web + debugging.
	// For native apps, CORS is generally not enforced, but supporting it keeps `web` tooling working.
	allowOrigin = strings.TrimSpace(allowOrigin)
	if allowOrigin == "" {
		allowOrigin = "*"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

