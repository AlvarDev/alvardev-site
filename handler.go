package main

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

var (
	templates = template.Must(template.New("").
		Funcs(template.FuncMap{}).ParseGlob("templates/*.html"))
)

type ctxKeyLog struct{}

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {

	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		//		"currencies":        currencies,
	}); err != nil {
		log.Fatal().Err(err).Msg("Can’t start service")
	}
}
