package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	InProduction  bool
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
}
