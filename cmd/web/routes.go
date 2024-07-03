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
	router.Use(SessionLoad)
	router.Use(NoSurf)

	// file server
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static", fs))

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	router.Get("/generals-quarters", handlers.Repo.Generals)
	router.Get("/search-availability", handlers.Repo.SearchAvailability)
	router.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	router.Get("/search-availability-json", handlers.Repo.AvailabilityJSON)
	router.Get("/majors-suite", handlers.Repo.Majors)
	router.Get("/make-reservation", handlers.Repo.MakeReservation)
	router.Get("/contact", handlers.Repo.Contact)
	return router
}
