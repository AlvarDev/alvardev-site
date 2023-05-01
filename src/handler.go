package main

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

var (
	tmpl *template.Template
	data templateData
)

// homHandler responds to requests by rendering an HTML page.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data = templateData{
		BlogActive: "is-active",
	}
	if err := tmpl.ExecuteTemplate(w, "home", data); err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("template.Execute: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

// aboutHandler responds to request by rendering an About HTML page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data = templateData{
		AboutActive: "is-active",
	}
	if err := tmpl.ExecuteTemplate(w, "about", data); err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("tempalte.Execute: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

// portafolioHandler responds to request by rendering a Portafolio HTML page
func portafolioHandler(w http.ResponseWriter, r *http.Request) {
	data = templateData{
		PortafolioActive: "is-active",
	}
	if err := tmpl.ExecuteTemplate(w, "portafolio", data); err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("tempalte.Execute: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
