// Package graph ADT graf dengan representasi adjacency matrix
package graph

import (
	"chizu-ru/node"
	"fmt"
	"math"
)

type vertex = node.Node

// graph ADT graf dengan adjacency matrix
type Graph struct {
	adjMatrix map[*vertex]map[*vertex]float64
}

// New fungsi untuk membuat sebuah graf baru
func New() *Graph {
	g := new(Graph)
	g.adjMatrix = make(map[*vertex]map[*vertex]float64)

	return g
}

// AddVertex fungsi untuk menambahkan sebuah sudut ke graf
// Jika sudut sudah ada, tidak akan dilakukan apa-apa
func (g *Graph) AddVertex(v *vertex) {
	if _, exists := g.adjMatrix[v]; exists {
		return
	}

	g.adjMatrix[v] = make(map[*vertex]float64)
	for k := range g.adjMatrix {
		g.adjMatrix[v][k] = math.Inf(1)
		g.adjMatrix[k][v] = math.Inf(1)
	}
	g.adjMatrix[v][v] = 0
}

// AddEdge fungsi untuk menambahkan sisi dari sudut src ke sudut dst
// dengan bobot weight
func (g *Graph) AddEdge(src *vertex, dst *vertex, weight float64) {
	g.AddVertex(src)
	g.AddVertex(dst)

	g.adjMatrix[src][dst] = weight
	g.adjMatrix[dst][src] = weight
}

func (g *Graph) Print() {
	for k, v := range g.adjMatrix {
		fmt.Print(*k)
		fmt.Println(":")
		for k2, v2 := range v {
			fmt.Printf("\t%s\t%f\n", k2.GetInfo(), v2)
		}
	}
}
