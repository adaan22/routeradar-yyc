# routeradar-yyc

This repo contains:

- `frontend/`: React Native app built with Expo (TypeScript)
- `backend/`: Go backend (health endpoint + API skeleton)

## Frontend (Expo Go)

```bash
cd frontend
npm install
npm run start
```

If you run on a physical device (Expo Go), set the backend base URL in `frontend/src/api/client.ts` (`API_BASE_URL`).
For now it defaults to `http://localhost:8080`.

## Backend (Go)

Go isn’t installed in this environment, but you can run it locally once you install Go:

```bash
cd backend
go run ./cmd/server -addr :8080
```

Test:

```bash
curl http://localhost:8080/health
```
