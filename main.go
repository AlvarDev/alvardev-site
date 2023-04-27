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
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const (
	port = "8080"
)

type frontendServer struct{}

func main() {
	svc := new(frontendServer)

	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("8080")
	}

	addr := os.Getenv("LISTEN_ADDR")

	r := mux.NewRouter()
	r.HandleFunc("/", svc.homeHandler).Methods(http.MethodGet, http.MethodHead)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/robots.txt", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "User-agent: *\nDisallow: /") })
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, "ok") })

	var handler http.Handler = r
	log.Info().Msg("starting server on " + addr + ":" + srvPort)

	if err := http.ListenAndServe(addr+":"+srvPort, handler); err != nil {
		log.Fatal().Err(err).Msg("Can’t start service")
	}
	//	log.Fatal()(http.ListenAndServe(addr+":"+srvPort, handler))
}
