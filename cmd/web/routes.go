package main

import (
	"net/http"

	"github.com/aiym182/news/internal/config"
	"github.com/aiym182/news/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	cors := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	mux.Use(cors)
	mux.Use(middleware.Recoverer)
	mux.Use(Nosurf)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/en", handlers.Repo.HomeEng)
	mux.Route("/en/pages", func(r chi.Router) {
		r.Get("/{num}", handlers.Repo.PagesEng)
	})
	mux.Route("/pages", func(r chi.Router) {
		r.Get("/{num}", handlers.Repo.Pages)
	})
	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
