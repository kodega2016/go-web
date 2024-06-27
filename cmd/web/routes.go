package main

import (
	"booking_app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Get("/about", handlers.About)
	router.Get("/", handlers.Home)
	return router
}
