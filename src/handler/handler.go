package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to Chizu-ru"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}
