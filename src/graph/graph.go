// Package graph ADT graf dengan representasi adjacency matrix
package graph

import (
	"chizu-ru/node"
	"chizu-ru/prioqueue"
	"errors"
	"fmt"
)

type vertex = node.Node

// graph ADT graf dengan adjacency matrix
type Graph struct {
	adjMatrix map[*vertex]map[*vertex]bool
}

// New fungsi untuk membuat sebuah graf baru
func New() *Graph {
	g := new(Graph)
	g.adjMatrix = make(map[*vertex]map[*vertex]bool)

	return g
}

// AddVertex fungsi untuk menambahkan sebuah sudut ke graf
// Jika sudut sudah ada, tidak akan dilakukan apa-apa
func (g *Graph) AddVertex(v *vertex) {
	if _, exists := g.adjMatrix[v]; exists {
		return
	}

	g.adjMatrix[v] = make(map[*vertex]bool)
	for k := range g.adjMatrix {
		g.adjMatrix[v][k] = false
		g.adjMatrix[k][v] = false
	}
	g.adjMatrix[v][v] = false
}

// AddEdge fungsi untuk menambahkan sisi dari sudut src ke sudut dst
// dengan bobot weight
func (g *Graph) AddEdge(src *vertex, dst *vertex) {
	g.AddVertex(src)
	g.AddVertex(dst)

	g.adjMatrix[src][dst] = true
	g.adjMatrix[dst][src] = true
}

func (g *Graph) Print() {
	for k, v := range g.adjMatrix {
		fmt.Print(*k)
		fmt.Println(":")
		for k2, v2 := range v {
			fmt.Printf("\t%s\t%t\n", k2.GetInfo(), v2)
		}
	}
}

func (g *Graph) AStar(src string, dest string) ([]*vertex, float64, error) {
	// graf baru yg nyimpen jarak spherical antar-sudut
	var srcVert, destVert *vertex
	srcVert = nil
	destVert = nil
	ret := make([]*vertex, 0)

	for k := range g.adjMatrix {
		if k.GetInfo() == src {
			srcVert = k
		} else if k.GetInfo() == dest {
			destVert = k
		}
	}

	if srcVert == nil || destVert == nil {
		return nil, 0, errors.New("sudut tujuan atau asal belum ada")
	}

	hasVisited := make(map[*vertex]bool)
	for k := range g.adjMatrix {
		hasVisited[k] = false
	}

	// nyari tetangga
	var fn, gn, hn float64
	curVert := srcVert
	for curVert != destVert {
		hasVisited[curVert] = true
		ret = append(ret, curVert)

		pq := prioqueue.New()
		for k := range g.adjMatrix[srcVert] {
			if !hasVisited[k] && g.adjMatrix[curVert][k] {
				if k != destVert {
					gn = node.Distance(srcVert, curVert)
					hn = node.Distance(curVert, destVert)
					fn = gn + hn
					pq.Enqueue(k, fn)
				} else { // hacky
					pq.Enqueue(k, -1)
				}
			}
		}

		nextVert, err := pq.Dequeue()
		if err != nil {
			return nil, 0, err
		}

		curVert = nextVert
	}

	ret = append(ret, destVert)
	return ret, fn, nil
}
