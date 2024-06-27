package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) error {
	// create template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("unable to create template cache:", err)
	}
	// get the template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("unable to find the template")
	}
	// render the template with the data
	buff := new(bytes.Buffer)
	t.Execute(w, nil)

	_, err = buff.WriteTo(w)
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
