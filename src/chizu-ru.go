package main

import (
	"chizu-ru/handler"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/AStar", handler.AStar)
	mux.HandleFunc("/Program", handler.ProgramHandler)

	handler := cors.Default().Handler(mux)
	log.Println("Starting web port on 8080")
	werr := http.ListenAndServe(":8080", handler)
	log.Fatal(werr)
}
