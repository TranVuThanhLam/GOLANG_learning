package handlers

import (
	"net/http"

	"modern_web.com/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "test.page.html")
}

// func About(w http.ResponseWriter, r *http.Request) {
// 	render.RenderTemplate(w, "about.page.tmpl")
// }
