package main

import (
	"booking_app/pkg/config"
	"booking_app/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	// use middleware
	// router.Use(WriteToConsole)
	router.Use(middleware.Logger)

	router.Get("/about", handlers.About)
	router.Get("/", handlers.Home)
	router.Get("/search", handlers.SearchAvailability)
	return router
}
