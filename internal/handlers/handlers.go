package handlers

import (
	"fmt"
	"net/http"

	"github.com/aiym182/news/internal/config"
	"github.com/aiym182/news/internal/models"
	"github.com/aiym182/news/internal/render"
	"github.com/go-chi/chi"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {

	render.Template(rw, r, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) HomeEng(rw http.ResponseWriter, r *http.Request) {
	render.Template(rw, r, "en.home.page.gohtml", &models.TemplateData{})
}
func (m *Repository) PagesEng(rw http.ResponseWriter, r *http.Request) {
	pageNum := chi.URLParam(r, "num")
	err := render.Template(rw, r, fmt.Sprintf("en.page-%s.page.gohtml", pageNum), &models.TemplateData{})
	if err != nil {
		rw.Write([]byte(fmt.Sprintln("Error loading pages")))
	}
}
func (m *Repository) Pages(rw http.ResponseWriter, r *http.Request) {

	pageNum := chi.URLParam(r, "num")

	err := render.Template(rw, r, fmt.Sprintf("page-%s.page.gohtml", pageNum), &models.TemplateData{})

	if err != nil {
		rw.Write([]byte(fmt.Sprintln("Error loading pages")))
	}
}
