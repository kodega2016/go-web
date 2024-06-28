package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application wide configurations
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	Session       *scs.SessionManager
}
