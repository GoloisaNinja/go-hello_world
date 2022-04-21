package handlers

import (
	"github.com/GoloisaNinja/go-hello_world/pkg/config"
	"github.com/GoloisaNinja/go-hello_world/pkg/models"
	"github.com/GoloisaNinja/go-hello_world/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (e *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	e.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (e *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again User!"

	remoteIP := e.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(
		w, "about.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
		},
	)
}
