package main

import (
	"context"
	"encoding/json"
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

func genaiBrHandler(w http.ResponseWriter, r *http.Request) {
	genaiBR := "https://docs.google.com/presentation/d/e/2PACX-1vQQLtEcDbQ7ZjNOmEyJDKf-NwRT5ymrPJ8IUHmf9eufaqgoFdYg6ZuktBdNbjMSl29mlS4gmM7Iz-7w/pub?start=false&loop=false&delayms=3000"
	http.Redirect(w, r, genaiBR, http.StatusSeeOther)
}

func genaiEsHandler(w http.ResponseWriter, r *http.Request) {
	genaiES := "https://docs.google.com/presentation/d/e/2PACX-1vRFI80CvjOjVf_ZxVNTHhOjtYKhp9FrxxkbVqBgtYiNNBf5Xz1yEYx2vhwIISfxPSxKwbTj3USg7uyc/pub?start=false&loop=false&delayms=3000"
	http.Redirect(w, r, genaiES, http.StatusSeeOther)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	ctx := context.Background()

	response := Response{}
	response.Posts = getFirestorePosts(ctx)
	responseJSON, err := json.Marshal(response)

	if err != nil {
		log.Printf("Failed to parse json: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

// Response definition for response API
type Response struct {
	Posts []Post `json:"posts"`
}
