package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deljumojtaba/golang-bookings/pkg/handlers"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	server_port := os.Getenv("PORT")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Server started at %s", server_port)
	http.ListenAndServe(":"+server_port, nil)
}
