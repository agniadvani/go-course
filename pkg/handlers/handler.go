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
	app *config.AppConfig
}

//Repository init function
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

//Handler init function
func NewHandler(r *Repository) {
	Repo = r
}

//Hnadler for "/home"
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.app.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//Handler for "/about"
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := map[string]string{
		"test": "Data from render.",
	}
	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
