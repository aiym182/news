package render

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/aiym182/news/internal/config"
	"github.com/aiym182/news/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var fs string = "../../template"
var functions = template.FuncMap{}

func NewTemplates(a *config.AppConfig) {

	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)

	return td
}

func Template(rw http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {

	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		tc = app.TemplateCache

	} else {
		tc, err = CTC()
		if err != nil {
			return err
		}
	}

	t, ok := tc[tmpl]

	if !ok {
		return errors.New("error finding template")
	}

	td = AddDefaultData(td, r)

	err = t.Execute(rw, td)

	if err != nil {
		return err
	}
	return nil
}

// create a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.gohtml", fs))

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.gohtml", fs))

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob(fmt.Sprintf("%s/*.layout.gohtml", fs))

			if err != nil {

				return myCache, nil
			}

			myCache[name] = ts

		}

	}

	return myCache, nil
}

func CTC() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*page.gohtml", fs))

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		myCache[name] = ts

	}

	return myCache, nil

}
