package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var guitars = map[string]string{
	"Fender": "$1000",
	"Gibson": "$2000",
	"Ibanez": "$3000",
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body><h1>John's Guitar Store</h1>"))
	w.Write([]byte("<h2><i>- Currently on Sale:</i></h2><ul>"))
	for k, v := range guitars {
		w.Write([]byte(fmt.Sprintf("<li>%s: %s</li>", k, v)))
	}
	w.Write([]byte("</ul></body></html>"))
}
