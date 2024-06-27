package config

import "text/template"

// AppConfig holds the application wide configurations
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
}
