package main

import (
	"chizu-ru/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/AStar", handler.AStar)
	mux.HandleFunc("/Program", handler.ProgramHandler)

	log.Println("Starting web port on 8080")
	werr := http.ListenAndServe(":8080", mux)
	log.Fatal(werr)
}
