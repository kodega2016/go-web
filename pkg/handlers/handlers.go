package handlers

import (
	"booking_app/pkg/config"
	"booking_app/pkg/models"
	"booking_app/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository holds the app config
type Repository struct {
	app *config.AppConfig
}

// NewRepository create a new instance of the repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

// NewHandler set the repo to the handler
func NewHandler(r *Repository) {
	Repo = r
}

func (re Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (re Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"name": "Khadga Bahadur Shrestha",
	}
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (re Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}
