package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/xman0142/goWebCourse/pkg/config"
	"github.com/xman0142/goWebCourse/pkg/handlers"
	"github.com/xman0142/goWebCourse/pkg/render"
)

const portNumber = ":8080"

// 1) A request comes in (a request for a web page)
// The request is captured by a route (in routes.go)
// If a matching route is found, it's handed to the handler specified by that route

// 2)  The template is loaded from disk into memory.

// 3) The template's parent (layout) and any partials are loaded from disk into memory

// 4) Logic, if any, is applied to the request in the handler (like a database lookup, or whatever)

// 5) The template is rendered, and data is passed to it, if necesary

// 6) The rendered template is handed to the responsewriter, and then sent to the end user's web browser.

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")

	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// //Routes
// http.HandleFunc("/", handlers.Repo.Home)
// http.HandleFunc("/about", handlers.Repo.About)
// http.HandleFunc("/teaching", handlers.Repo.Teaching)
// http.HandleFunc("/entreneurship", handlers.Repo.Entreneurship)
