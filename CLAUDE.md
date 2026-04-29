# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Structure

- `frontend/` — React Native app (Expo, TypeScript)
- `backend/` — Go HTTP server (stdlib only, no third-party deps)

## Commands

### Frontend

```bash
cd frontend
npm install        # install deps
npm run start      # start Expo dev server (scan QR with Expo Go)
npm run ios        # iOS simulator
npm run android    # Android emulator
npm run web        # web browser
```

### Backend

```bash
cd backend
go run ./cmd/server -addr :8080                          # start server
go run ./cmd/server -addr :8080 -cors-allow-origin '*'  # explicit CORS origin
curl http://localhost:8080/health                        # smoke test
```

## Architecture

### Backend

Entry point: `backend/cmd/server/main.go` — parses `-addr` and `-cors-allow-origin` flags, builds the router, and runs the server with graceful shutdown on SIGINT/SIGTERM.

- `internal/api/routes.go` — registers all HTTP handlers on `http.ServeMux`, wraps the mux with CORS middleware
- `internal/api/health.go` — `GET /health` returns `{"status":"ok"}`; `HEAD /health` returns headers only
- `internal/api/cors.go` — CORS middleware wrapping; native apps don't enforce CORS but it's needed for `expo web`
- `internal/server/server.go` — thin wrapper around `http.Server` with configurable timeouts and graceful shutdown

New API endpoints go under `/api/v1/` in `routes.go`.

### Frontend

Entry point: `frontend/index.ts` → `App.tsx`. Currently a single-screen app that calls the backend health endpoint on mount and displays the result.

- `frontend/src/api/client.ts` — all backend communication; `API_BASE_URL` defaults to `http://localhost:8080`. Change this constant when running on a physical device (Expo Go) to point at your machine's LAN IP.

The Expo app has `newArchEnabled: true` (React Native New Architecture).
