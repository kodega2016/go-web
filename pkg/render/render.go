package render

import (
	"booking_app/pkg/config"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config.AppConfig

// NewTemplates set the app config to the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) error {
	tc := map[string]*template.Template{}

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
	t.Execute(w, nil)

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
