package main

import (
	"chizu-ru/graph"
	"chizu-ru/vertex"
)

func main() {
	n1 := vertex.New("hadeh", 1, 1)
	n2 := vertex.New("mabar", 0, -100)
	n3 := vertex.New("anjay", 69, 420)

	g := graph.New()
	g.AddEdge(n1, n2, 1)
	g.AddEdge(n3, n1, 2)
	g.Print()
}
