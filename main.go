package main

import (
	"log"
	"net/http"
	"os"

	"stockpile-alert-backend/src/routes"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // Default to development
	}
	err := godotenv.Load("env/" + env + ".env")
	if err != nil {
		log.Fatalf("Error loading %s environment file: %v", env, err)
	}
}

func main() {
	// Initialize routes
	router := routes.InitializeRoutes()

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
