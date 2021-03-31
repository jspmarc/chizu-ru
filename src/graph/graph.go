// Package graph ADT graf dengan representasi adjacency matrix
package graph

import (
	. "chizu-ru/vertex"
	"fmt"
	"math"
)

// graph ADT graf dengan adjacency matrix
// Disertai slice of pointer to vertex untuk mendapatkan list vertex
// yang dimiliki graf
type graph struct {
	adjMatrix map[*Vertex]map[*Vertex]float64
	vertList   []*Vertex
}

// New fungsi untuk membuat sebuah graf baru
func New() *graph {
	g := new(graph)
	g.adjMatrix = make(map[*Vertex]map[*Vertex]float64)
	g.vertList = make([]*Vertex, 0)

	return g
}

// AddVertex fungsi untuk menambahkan sebuah sudut ke graf
// Jika graf sudah ada, tidak akan dilakukan apa-apa
func (g *graph) AddVertex(v *Vertex) {
	if _, exists := g.adjMatrix[v]; exists {
		return
	}

	g.vertList = append(g.vertList, v)
	g.adjMatrix[v] = make(map[*Vertex]float64)
	for k := range g.adjMatrix {
		g.adjMatrix[v][k] = math.Inf(1)
		g.adjMatrix[k][v] = math.Inf(1)
	}
	g.adjMatrix[v][v] = 0
}

// AddEdge fungsi untuk menambahkan sisi dari sudut src ke sudut dst
// dengan bobot weight
func (g *graph) AddEdge(src *Vertex, dst *Vertex, weight float64) {
	g.AddVertex(src)
	g.AddVertex(dst)

	g.adjMatrix[src][dst] = weight
}

func (g *graph) Print() {
	for k, v := range g.adjMatrix {
		fmt.Print(*k)
		fmt.Println(":")
		for k2, v2 := range v {
			fmt.Printf("\t%s\t%f\n", k2.GetInfo(), v2)
		}
	}
}
