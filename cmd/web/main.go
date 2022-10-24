package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xman0142/goWebCourse/pkg/config"
	"github.com/xman0142/goWebCourse/pkg/handlers"
	"github.com/xman0142/goWebCourse/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")

	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/teaching", handlers.Teaching)
	http.HandleFunc("/entreneurship", handlers.Entreneurship)

	fmt.Println(fmt.Printf("Starting Application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
