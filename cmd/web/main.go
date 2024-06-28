package main

import (
	"booking_app/pkg/config"
	"booking_app/pkg/handlers"
	"booking_app/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fatih/color"
)

const portNumber = ":8080"

var session *scs.SessionManager

func main() {
	// initialize the app
	var app config.AppConfig
	app.InProduction = false
	app.UseCache = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// create template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("unable to create template cache:", err)
	}
	app.TemplateCache = tc

	// sets app config to the render package
	render.NewTemplates(&app)

	// create new repo for the handlers
	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	// server configuration
	srv := http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	c := color.New(color.FgGreen).Add(color.Underline)
	c.Printf("application server is running on %s.\n", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
