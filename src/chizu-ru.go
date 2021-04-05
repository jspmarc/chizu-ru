package main

import (
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
	mux.HandleFunc("/upload", handler.Upload)

	handler := cors.Default().Handler(mux)
	log.Println("Starting web port on 8080")
	werr := http.ListenAndServe(":8080", handler)
	log.Fatal(werr)
}
