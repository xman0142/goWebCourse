package handlers

import (
	"net/http"

	"github.com/xman0142/goWebCourse/pkg/config"
	"github.com/xman0142/goWebCourse/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is about rosie handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// Teaching is about Rosie Teaching experience
func (m *Repository) Teaching(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "teaching.page.tmpl")
}

// Entrepenship is entrpenur handler
func (m *Repository) Entreneurship(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "entreneurship.page.tmpl")

}
