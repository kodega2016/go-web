package render

import (
	"booking_app/pkg/config"
	"booking_app/pkg/models"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// NewTemplates set the app config to the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultValue(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = "this is just a demo message"
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get the template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("unable to find the template")
	}
	// render the template with the data
	buff := new(bytes.Buffer)
	td = AddDefaultValue(td, r)
	t.Execute(w, td)

	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("error writing to the template:", err)
		return err
	}
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tc := map[string]*template.Template{}
	var err error

	// get all the pages from the ./templates/*.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	// iterate over the pages and create the new template set
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tc, err
		}

		// get all the layout files from the ./templates directory
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tc, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tc, err
			}
		}
		tc[name] = ts
	}
	return tc, err
}
