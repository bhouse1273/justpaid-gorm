# JustPaid Echo API (GORM)
# justpaid-gorm

GORM-backed Go API and worker for the JustPaid MySQL schema. The project exposes a small Echo HTTP API and a worker that uses stored procedures in the attached schema.

Repository layout
-----------------
- `cmd/api` — Echo v4 HTTP API server
- `cmd/worker` — background worker (claims and completes actions)
- `cmd/justpaid-demo` — small demo runner (inserts sample data, calls procs)
- `internal/db`, `internal/models`, `internal/handlers`, `internal/config` — core packages

Prerequisites
-------------
- Go 1.24+
- MySQL 8 (the provided SQL uses DATETIME(6) and stored procedures)

Database
--------
Load the full schema (example):

```sh
mysql -u <user> -p < /path/to/justpaid_full_datetime.sql
```

Ensure your DSN includes `parseTime=true` so `time.Time` fields work correctly.

Environment
-----------
Create a `.env` file or export variables. Example `.env`:

```env
MYSQL_DSN=user:pass@tcp(localhost:3306)/justpaid?parseTime=true&loc=Local
PORT=8080
CORS_ORIGINS=http://localhost:3000
JWT_SECRET=some-secret-if-you-want-auth
WORKER_ENABLED=1
WORKER_BATCH=200
WORKER_TICK_SECONDS=30
```

Build & run
-----------
Tidy deps and build:

```sh
go mod tidy
go build ./...
```

Run the API server (reads `.env`):

```sh
go run ./cmd/api
```

Run the worker:

```sh
go run ./cmd/worker
```

Run the demo runner:

```sh
go run ./cmd/justpaid-demo
```

API endpoints
-------------
- `POST /api/campaigns/:id/materialize` — materialize actions for a campaign (calls stored proc)
- `POST /api/actions/claim` — claim due actions (body: `{ "batch": <n> }`) — returns rows
- `POST /api/actions/:id/complete` — mark an action complete (body: `{ "success": true, "result": {...} }`)

Notes for clients (React)
------------------------
- Money values (`decimal.Decimal`) are JSON-encoded as strings; use a big-number lib or parse as string.
- DATETIME(6) fields are returned as RFC3339/ISO strings; format in UI with date-fns/dayjs.
- Enable CORS via `CORS_ORIGINS` for the React dev origin.

VS Code and gopls
-----------------
If VS Code shows import errors after module changes, restart the Go language server:

- Command Palette -> "Go: Restart Language Server"
- Or: Developer: Reload Window

CI suggestion
-------------
Add a GitHub Actions workflow that runs `go mod tidy`, `go vet`, `go test ./...` and `go build ./...` on push.

Tests
-----
There are currently no unit tests. Recommended next steps:

- Add handler unit tests using `httptest` and a mocked GORM DB.
- Add integration tests against a disposable MySQL instance (Docker) for stored-proc behavior.

License
-------
Add a LICENSE file when you choose a license for this code.
