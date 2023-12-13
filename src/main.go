// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	// Prepare template for execution.
	tmpl = template.Must(template.New("").Funcs(template.FuncMap{}).ParseGlob("templates/*.html"))

	// Define HTTP server.
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/genaibr", genaiBrHandler)
	http.HandleFunc("/genaies", genaiEsHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Can’t start service")
	}
}
