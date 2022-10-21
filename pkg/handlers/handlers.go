package handlers

import (
	"net/http"
)

// Home is homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl")
}

// About is about rosie handler
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl")
}

// Teaching is about Rosie Teaching experience
func Teaching(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "teaching.page.tmpl")
}

// Entrepenship is entrpenur handler
func Entreneurship(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "entreneurship.page.tmpl")

}
