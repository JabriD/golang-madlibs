package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	if friendsname := r.FormValue("friendsname"); friendsname != "" {
		log.Printf("Friend's Name: %s", r.FormValue("friendsname"))
	}
	if adjective := r.FormValue("adjective"); adjective != "" {
		log.Printf("Adjective: %s", r.FormValue("adjective"))
	}
	if noun := r.FormValue("noun"); noun != "" {
		log.Printf("Noun: %s", r.FormValue("noun"))
	}
	if place := r.FormValue("place"); place != "" {
		log.Printf("Place: %s", r.FormValue("place"))
	}
	if verb := r.FormValue("verb"); verb != "" {
		log.Printf("Verb: %s", r.FormValue("verb"))
	}
	if noun2 := r.FormValue("noun2"); noun2 != "" {
		log.Printf("Noun: %s", r.FormValue("noun2"))
	}
	if yourname := r.FormValue("yourname"); yourname != "" {
		log.Printf("Your Name: %s", r.FormValue("yourname"))
	}

	t.templ.Execute(w, r)
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/greet", &templateHandler{filename: "greeting.html"})
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
