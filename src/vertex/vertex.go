// Package vertex ADT vertex untuk graf
package vertex

import (
	"math"
)

type coor struct {
	x float64
	y float64
}

// Vertex ADT untuk sudut graf
type Vertex struct {
	info string
	pos  coor
}

// New fungsi untuk membuat sebuah vertex baru dengan
func New(info string, x float64, y float64) *Vertex {
	node := new(Vertex)
	node.info = info
	node.pos.x = x
	node.pos.y = y

	return node
}

// CartesianDistance fungsi untuk menghitung jarak cartesian antara dua buah
// sudut
func CartesianDistance(v1 *Vertex, v2 *Vertex) float64 {
	xDist := v1.pos.x - v2.pos.x
	yDist := v1.pos.y - v2.pos.y

	return math.Sqrt(math.Pow(xDist, 2) + math.Pow(yDist, 2))
}

func (v *Vertex) GetInfo() string {
	return v.info
}
