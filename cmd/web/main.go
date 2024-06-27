package main

import (
	"github.com/fatih/color"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// server configuration
	srv := http.Server{
		Addr:    portNumber,
		Handler: Routes(),
	}

	c := color.New(color.FgGreen).Add(color.Underline)
	c.Printf("application server is running on %s.\n", portNumber)
	srv.ListenAndServe()
}
