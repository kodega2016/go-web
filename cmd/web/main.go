package main

import (
	"booking_app/pkg/config"
	"booking_app/pkg/render"
	"log"
	"net/http"

	"github.com/fatih/color"
)

const portNumber = ":8080"

func main() {
	// initialize the app
	var app config.AppConfig
	app.InProduction = false
	app.UseCache = false

	// create template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("unable to create template cache:", err)
	}
	app.TemplateCache = tc

	// sets app config to the render package
	render.NewTemplates(&app)

	// server configuration
	srv := http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	c := color.New(color.FgGreen).Add(color.Underline)
	c.Printf("application server is running on %s.\n", portNumber)
	srv.ListenAndServe()
}
