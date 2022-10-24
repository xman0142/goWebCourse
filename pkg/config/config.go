package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
}
