# <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" height="24"> Go Template

Minimal Go HTTP server with routing, middleware, and static file support using only the standard library.

## Local Development

Build and run:
```bash
go build -o server ./cmd/server && ./server
```

Or run directly:
```bash
go run ./cmd/server
```

Server runs on http://localhost:3000

## File Structure

```
├── cmd/
│   └── server/
│       └── main.go         # Application entry point
├── internal/
│   ├── handler/            # HTTP handlers
│   │   ├── health.go
│   │   ├── user.go
│   │   └── json.go
│   └── middleware/         # HTTP middleware
│       └── logging.go
├── public/                 # Static files
├── go.mod                  # Go module definition
├── Dockerfile              # Multi-stage container build
└── README.md
```

## API Endpoints

- `GET /` - Serves static files from public/
- `GET /health` - Returns `{"status": "ok"}`
- `GET /users` - Returns `{"message": "respond with a resource"}`

---

## <img src="https://github.com/sugasrc.png" height="20"> Deploying with Suga

1. Click **Use this template** to create your own repository

2. Push changes - GitHub Actions builds and pushes to GHCR automatically

3. In the Suga dashboard, update the **Image URI** for your service:
   ```
   ghcr.io/YOUR_USERNAME/YOUR_REPO:latest
   ```

4. Ensure your GHCR package is public
