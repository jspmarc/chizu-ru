package main

import (
	//"chizu-ru/parser"
	//"fmt"
	"chizu-ru/handler"
	"log"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./views"))
	mux.Handle("/", http.StripPrefix("/", fs))
	mux.HandleFunc("/AStar", handler.AStar)

	handler := cors.Default().Handler(mux)
	log.Println("Starting web port on 8080")
	werr := http.ListenAndServe(":8080", handler)
	log.Fatal(werr)

	//g, _ := parser.ParseFile("../test/2.txt")
	//fmt.Println(g)
	//fmt.Println(g.AStar("Arad", "Iasi"));
}
