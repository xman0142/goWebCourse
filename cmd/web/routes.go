package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/xman0142/goWebCourse/pkg/config"
	"github.com/xman0142/goWebCourse/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/entreneurship", http.HandlerFunc(handlers.Repo.Entreneurship))
	mux.Get("/teaching", http.HandlerFunc(handlers.Repo.Teaching))

	return mux
}
