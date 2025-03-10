package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/thanhlam/bookings/pkg/config"
	"github.com/thanhlam/bookings/pkg/models"
)

// filepath.Glob()
// filepath.Base()

var tc = make(map[string]*template.Template)

var app *config.AppConfig

// Newtemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	// Ensure td is not nil
	if td == nil {
		td = &models.TemplateData{}
	}

	err := t.Execute(buf, td) // Pass td instead of nil
	if err != nil {
		log.Fatal(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files name *.page from ./template
	pages, err := filepath.Glob("./templates/*.page")
	if err != nil {
		return nil, err
	}

	// range through all files endign with *.page
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
