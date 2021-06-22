package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

//Tools to be used by a variable
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
