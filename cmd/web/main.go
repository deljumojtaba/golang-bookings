package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deljumojtaba/golang-bookings/pkg/config"
	"github.com/deljumojtaba/golang-bookings/pkg/handlers"
	"github.com/deljumojtaba/golang-bookings/pkg/render"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	server_port := os.Getenv("PORT")

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Server started at %s", server_port)
	http.ListenAndServe(":"+server_port, nil)
}
