package handler

import (
	"chizu-ru/parser"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path"
)

type aStarRes struct {
	Distance float64
	Nodes    []resNode
}

type resNode struct {
	Name      string
	Latitude  float64
	Longitude float64
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to Chizu-ru"))
	tmpl, err := template.ParseFiles(path.Join("views", "Home.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func ProgramHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "Program.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AStar format query: AStar?filePath=[filePath]&src=[src]&dest=[dest]
func AStar(w http.ResponseWriter, r *http.Request) {
	s := r.URL
	u, err := url.Parse(s.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m, _ := url.ParseQuery(u.RawQuery)

	if len(m["filePath"]) == 0 || len(m["src"]) == 0 || len(m["dest"]) == 0 {
		http.Error(w, "get query salah, butuh `filePath`, `src`, dan `dest`", http.StatusBadRequest)
		return
	}

	g, err := parser.Parse(m["filePath"][0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nodes, jarak, _ := g.AStar(m["src"][0], m["dest"][0])
	res := new(aStarRes)
	for i := len(nodes) - 1; i > -1; i-- {
		res.Nodes = append(res.Nodes,
			resNode{Name: nodes[i].GetInfo(),
				Latitude:  nodes[i].GetLat(),
				Longitude: nodes[i].GetLong()},
		)
	}
	res.Distance = jarak
	jsRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsRes)
}
