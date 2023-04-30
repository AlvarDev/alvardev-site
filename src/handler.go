package main

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

var (
	tmpl *template.Template
)

// homHandler responds to requests by rendering an HTML page.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "home", nil); err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("template.Execute: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
