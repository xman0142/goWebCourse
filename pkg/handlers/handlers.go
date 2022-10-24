package handlers

import (
	"net/http"

	"github.com/xman0142/goWebCourse/pkg/render"
)

// Home is homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is about rosie handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// Teaching is about Rosie Teaching experience
func Teaching(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "teaching.page.tmpl")
}

// Entrepenship is entrpenur handler
func Entreneurship(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "entreneurship.page.tmpl")

}
