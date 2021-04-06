// Package graph ADT graf dengan representasi adjacency matrix
package graph

import (
	"chizu-ru/node"
	"chizu-ru/prioqueue"
	"errors"
	"fmt"
	//"time"
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
		return nil, 0, errors.New("EdgeNotFound")
	}

	prevVerts := make(map[*vertex]*vertex)
	vertCostFromSrc := make(map[*vertex]float64)
	toVisit := prioqueue.New()

	vertCostFromSrc[srcVert] = 0
	prevVerts[srcVert] = nil
	toVisit.Enqueue(srcVert, 0)

	// nyari tetangga
	var fn, gn, hn float64
	fn = 0
	gn = 0
	hn = 0
	var curVert *vertex = nil

	for !toVisit.IsEmpty() {
		var err error
		curVert, err = toVisit.Dequeue()
		if err != nil {
			return nil, 0, errors.New("NoConnectionToDest")
		}

		for adj, isAdj := range g.adjMatrix[curVert] {
			if isAdj {
				// hitung gn
				gn = vertCostFromSrc[curVert] + node.Distance(curVert, adj)
				// jadiin gn sbg cost dari source vertex ke k
				if cost, exists := vertCostFromSrc[adj]; !exists {
					vertCostFromSrc[adj] = gn
					prevVerts[adj] = curVert
				} else {
					if gn < cost {
						// update jarak dari source, kalo ternyata bisa lebih
						// kecil. Update juga parent vertex-nya
						vertCostFromSrc[adj] = gn
						prevVerts[adj] = curVert
					}
				}
				// kalo ketemu deestiation vertex, stop
				if adj == destVert {
					toVisit.Clear()
					break
				}
				// hitung hn
				hn = node.Distance(adj, destVert)
				// hitung fn
				fn = gn + hn

				// ga usah dimasukin kalo udah ada di queue dengan cost lebih
				// kecil atau jarak dari src lebih kecil dari gn yang dihitung
				if toVisit.ContainsLEQ(adj, fn) || gn > vertCostFromSrc[adj] {
					continue
				}

				toVisit.RemoveEl(adj)
				toVisit.Enqueue(adj, fn)
			}
		}
	}

	curVert = destVert
	for curVert != nil {
		prevVert := prevVerts[curVert]
		ret = append(ret, curVert)
		curVert = prevVert
	}

	// reverse result graph
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret, vertCostFromSrc[destVert], nil
}
