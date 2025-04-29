package main

import (
	"fmt"
	"github.com/MikkelAllison/rilla-backend/internal/auth"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default to 8080 if not specified
	}

	// Create a new router
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Public routes
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/signup", auth.SignUpHandler)
		r.Post("/login", auth.LoginHandler)
	})

	fmt.Println("Server starting on port", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
