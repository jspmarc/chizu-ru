package handler

import (
	"chizu-ru/graph"
	"chizu-ru/parser"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
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

func Upload(w http.ResponseWriter, r *http.Request) {
	log.Println("/upload endpoint hit")
	if r.Method != "POST" {
		http.Error(w, "Request must be a POST method", http.StatusBadRequest)
		log.Println("Request not POST")
		return
	}

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("uploaded", handler.Filename+".txt")
	if err != nil {
		log.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	fmt.Fprint(w, fileBytes)
	g, _ := parser.Parse("uploaded/" + handler.Filename + ".txt")
	g.Print()
	log.Println("File" + handler.Filename + ".txt uploaded")
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
