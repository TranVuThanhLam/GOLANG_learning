package main

import (
	"fmt"
	"log"
	"net/http"

	"modern_web.com/pkg/config"
	"modern_web.com/pkg/handlers"
	"modern_web.com/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("server is running on port%s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
