# JustPaid Echo API (GORM)

Lean Echo-based API & worker for your JustPaid MySQL schema.

## What's included
- Echo v4 HTTP server with CORS, logging, recover
- Optional JWT middleware (enable with `JWT_SECRET`)
- Endpoints:
  - `POST /api/campaigns/:id/materialize`
  - `POST /api/actions/claim` (body: `{ "batch": 200 }`)
  - `POST /api/actions/:id/complete` (body: `{ "success": true, "result": {...} }`)
- Background worker (`cmd/worker`) using your stored procedures
- GORM connection with singular table names
- Models for the Campaign flow (expand as needed)

## Setup
1. Ensure DB is created with `justpaid_full_datetime.sql`.
2. Copy `.env.example` to `.env` and set `MYSQL_DSN` (keep `parseTime=true`).
3. Run API:
```bash
go run ./cmd/api
```
4. Run worker (optional):
```bash
WORKER_ENABLED=1 go run ./cmd/worker
```

## Frontend example
```ts
await fetch('/api/campaigns/UUID/materialize', { method: 'POST', credentials: 'include' })
```

You can add more handlers and models as your UI grows. Echo keeps dependencies minimal and routing straightforward.
