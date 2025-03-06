package config

import (
	"html/template"
	"log"
)

// app hold the application configuration config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
