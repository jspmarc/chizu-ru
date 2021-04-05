package handler

import (
	"chizu-ru/graph"
	"chizu-ru/parser"
	"encoding/json"
	"html/template"
	"io"
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
	log.Println("/ hit")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
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

// AStar format query: AStar?filePath=[filePath]&src=[src]&dest=[dest]
func AStar(w http.ResponseWriter, r *http.Request) {
	log.Println("/AStar endpoint hit")
	u := r.URL
	m, _ := url.ParseQuery(u.RawQuery)

	if len(m["src"]) == 0 || len(m["dest"]) == 0 {
		http.Error(w, "src and dest not set. hint: /Astar?dest=Arad&src=Bucharest", http.StatusBadRequest)
		return
	}

	var g *graph.Graph
	var err error
	if len(m["filePath"]) != 0 || r.Method != "POST" {
		if len(m["filePath"]) == 0 {
			http.Error(w, "filePath query not set", http.StatusBadRequest)
		}
		g, err = parser.ParseFile(m["filePath"][0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else { // parse request body. Request body is the file's content
		var content []byte
		content, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		g, err = parser.Parse(string(content))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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
