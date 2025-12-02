# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/api

# Runtime stage
FROM scratch

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/public ./public

EXPOSE 3000

ENTRYPOINT ["./server"]
