package main

import (
	"chizu-ru/graph"
	"chizu-ru/node"
	"fmt"
	//"chizu-ru/osm"
)

func main() {
	n1 := node.New("hadeh", 1, 1)
	n2 := node.New("mabar", 0, -100)
	n3 := node.New("anjay", 69, 420)

	g := graph.New()
	g.AddEdge(n1, n2, 1)
	g.AddEdge(n3, n1, 2)
	g.Print()

	newYork := node.New("New York", 40.7128, 74.0060)
	london := node.New("London", 51.5074, 0.1278)

	fmt.Println(node.Distance(newYork, london))
	//osm.Search("name", "Institut Teknologi Bandung", "Jawa Barat", 25)
}
