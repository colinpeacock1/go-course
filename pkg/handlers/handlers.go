package handlers

import (
	"net/http"

	"github.com/colinpeacock1/go-course/pkg/render"
)

// Home - the homepage
func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl")
}

// About - the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl")
}
