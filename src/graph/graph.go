// Package graph ADT graf dengan representasi adjacency matrix
package graph

import (
	"chizu-ru/node"
	"chizu-ru/prioqueue"
	"errors"
	"fmt"
)

type vertex = node.Node

// Graph ADT graf dengan adjacency matrix
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

	prevVerts := make(map[*vertex]*vertex)
	vertCostFromSrc := make(map[*vertex]float64)
	vertCostFromSrc[srcVert] = 0
	prevVerts[srcVert] = nil

	// nyari tetangga
	var fn, gn, hn float64
	fn = 0
	gn = 0
	hn = 0
	var curVert *vertex = nil
	pq := prioqueue.New()
	pq.Enqueue(srcVert, 0)

	for !pq.IsEmpty() {
		var err error
		curVert, err = pq.Dequeue()
		if err != nil {
			return nil, 0, err
		}

		fmt.Println(*curVert)
		for k, v := range g.adjMatrix[curVert] {
			if v {
				if k == destVert {
					for !pq.IsEmpty() {
						pq.Dequeue()
					}
					break
				}
				// hitung gn
				gn = vertCostFromSrc[curVert] + node.Distance(curVert, k)
				// jadiin gn sbg cost dari source vertex ke k
				if _, exists := vertCostFromSrc[k]; !exists {
					vertCostFromSrc[k] = gn
					prevVerts[k] = curVert
				} else {
					if gn < vertCostFromSrc[k] {
						// update jarak dari source, kalo ternyata bisa lebih
						// kecil. Update juga parent vertex-nya
						vertCostFromSrc[k] = gn
						prevVerts[k] = curVert
					}
				}
				// hitung hn
				hn = node.Distance(k, destVert)
				// hitung fn
				fn = gn + hn

				fmt.Println("\t", k.GetInfo(), gn, "+", hn, "=", fn)
				pq.Enqueue(k, fn)
				tempQ := *pq
				fmt.Print("\t[")
				for !tempQ.IsEmpty() {
					e, _ := tempQ.Dequeue()
					fmt.Print(e.GetInfo(), " ")
				}
				fmt.Println("]")
			}
		}
	}

	curVert = destVert
	for curVert != nil {
		prevVert := prevVerts[curVert]
		if prevVert == nil {
			break
		}
		ret = append(ret, curVert)
		curVert = prevVert
	}
	return ret, vertCostFromSrc[destVert], nil
}
