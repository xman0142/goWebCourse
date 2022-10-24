package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/xman0142/goWebCourse/pkg/config"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Renders the Html Template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {

		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template chache")

	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the file paths named *.page.tmpl from ./templates folder, about, home, teaching etc
	//Glob goes to a certain location ("./templates/*.page.tmpl") and chooses all thee different files
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	// check for error in case templates folder doesnt exsist
	if err != nil {
		return myCache, err
	}

	// range through all files ending with the pattern of "*.page.tmpl"  * is the pattern of the name
	for _, page := range pages {
		// returns the end of page name aka About,
		name := filepath.Base(page)
		// add files that ts is Template set
		// Previous code: ts, err := template.New(name).ParseFiles(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// checks to see if there are any files that matche the pattern "./templates/*.layout.tmpl"

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// if there are any files that match "./templates/*.layout.tmpl" then read files with ./templates/*.layout.tmpl" and add to ts (template set)

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
