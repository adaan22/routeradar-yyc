# routeradar-yyc backend

## Prerequisites

- Go installed (this repo currently uses the standard library only, so there are no extra third-party dependencies yet).

## Run

```bash
cd backend
go run ./cmd/server -addr :8080
```

Health check:

```bash
curl http://localhost:8080/health
```

