package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/deljumojtaba/golang-bookings/internal/config"
	"github.com/deljumojtaba/golang-bookings/internal/handlers"
	"github.com/deljumojtaba/golang-bookings/internal/render"
	"github.com/joho/godotenv"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	server_port := os.Getenv("PORT")

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	log.Printf("Server started at %s", server_port)

	srv := &http.Server{
		Addr:    ":" + server_port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
