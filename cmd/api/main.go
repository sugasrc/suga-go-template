package main

import (
	"log"
	"net/http"
	"os"

	"suga-go-template/internal/handler"
	"suga-go-template/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /users", handler.Users)

	// Static file serving from public/ directory
	mux.Handle("/", http.FileServer(http.Dir("public")))

	// Apply middleware
	h := middleware.Logging(mux)

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal(err)
	}
}
