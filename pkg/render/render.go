package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Renders the Html Template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)

	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
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
		ts, err := template.New(name).ParseFiles(page)
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
