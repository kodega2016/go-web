package handlers

import (
	"booking_app/pkg/config"
	"booking_app/pkg/models"
	"booking_app/pkg/render"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

var Repo *Repository

// Repository holds the app config
type Repository struct {
	App *config.AppConfig
}

// NewRepository create a new instance of the repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler set the repo to the handler
func NewHandler(r *Repository) {
	Repo = r
}

func (re Repository) Home(w http.ResponseWriter, r *http.Request) {
	Repo.App.Session.Put(r.Context(), "ip_address", r.RemoteAddr)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (re Repository) About(w http.ResponseWriter, r *http.Request) {
	ip_address := re.App.Session.GetString(r.Context(), "ip_address")
	stringMap := map[string]string{
		"name":       "Khadga Bahadur Shrestha",
		"ip_address": ip_address,
	}
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (re Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (re Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	start_date := r.Form.Get("start")
	end_date := r.Form.Get("end")
	fmt.Fprintln(w, start_date, end_date)
}

func (re Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	res := jsonResponse{
		OK:      true,
		Message: "json data is stored successfully.",
	}
	out, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (re Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (re Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}
func (re Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (re Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
