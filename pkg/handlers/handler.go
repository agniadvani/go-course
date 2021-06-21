package handler

import (
	"net/http"

	"github.com/agniadvani/go-course/pkg/config"
	"github.com/agniadvani/go-course/pkg/models"
	render "github.com/agniadvani/go-course/pkg/renders"
)

//A variable used to reference outide the package scope values
var Repo *Repository

//Struct that holds the variable app, can be referenced outside the package scope using New Repo
type Repository struct {
	App *config.AppConfig
}

//Repository init function
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//Handler init function
func NewHandler(r *Repository) {
	Repo = r
}

//Hnadler for "/home"
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//Handler for "/about"
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := map[string]string{
		"test": "Data from render.",
	}
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
