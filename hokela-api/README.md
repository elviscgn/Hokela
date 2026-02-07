Small Go API using Gin. This README explains how to build and run the service locally.

Prerequisites
- Go (see `go.mod` for the required version) — e.g. Go 1.25

Quick start
1. Clone the repository and change into the project directory:

```bash
git clone <repo-url>
cd hokela-api
```

2. Download dependencies:

```bash
go mod tidy
```

3. Run the server (default port 8080):

```bash
go run ./cmd/api
```

Or build and run the binary:

```bash
go build -o bin/hokela-api ./cmd/api
./bin/hokela-api
```

Custom port
Set the `PORT` environment variable to change the listen port. Example:

```bash
PORT=3000 go run ./cmd/api
```

API endpoints
- Health check: `GET /api/ping` — returns a JSON status and timestamp.

Example request

```bash
curl http://localhost:8080/api/ping
# => {"status":"Healthy","timestamp":<unix_ts>}
```

Notes
- The project includes `github.com/joho/godotenv` in `go.mod` but loading a `.env` file in `cmd/api/main.go` is currently commented out. If you'd like to enable `.env` support, uncomment the `godotenv` import and load call in `cmd/api/main.go`.
- See [cmd/api/main.go](cmd/api/main.go) for startup logic and default port.

If you want, I can add a Makefile, Dockerfile, or example .env file next.

To run: `go run cmd/api/main.go`