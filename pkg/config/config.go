package config

import "text/template"

//Tools to be used by a variable
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
}
