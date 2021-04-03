// Package node ADT vertex untuk graf
package node

import (
	"math"
)

type gcs struct {
	latitude  float64 // N, satuan derajat
	longitude float64 // W, satuan derajat
}

// Node ADT untuk sudut graf
// gcs adalah global coordinate system yang menandakan longitude dan latitude
// sudut dengan satuan derajat
type Node struct {
	info string
	pos  gcs
}

// New fungsi untuk membuat sebuah vertex baru dengan
func New(info string, latitude float64, longitude float64) *Node {
	node := new(Node)
	node.info = info
	node.pos.longitude = longitude
	node.pos.latitude = latitude

	return node
}

func toRadian(deg float64) float64 {
	return deg * math.Pi / 180
}

func haversine(rad float64) float64 {
	return math.Pow(math.Sin(rad/2), 2)
}

// Distance fungsi untuk menghitung jarak antara dua buah sudut dengan
// menggunakan persamaan great-sphecial distance. Hasilnya dalam kilometer
func Distance(v1 *Node, v2 *Node) float64 {
	r := 6371.0 // jari-jari bumi
	v1lat := toRadian(v1.pos.latitude)
	v1lon := toRadian(v1.pos.longitude)
	v2lat := toRadian(v2.pos.latitude)
	v2lon := toRadian(v2.pos.longitude)

	h := haversine(v2lat-v1lat) +
		math.Cos(v1lat)*math.Cos(v2lat)*haversine(v2lon-v1lon)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func (v *Node) GetInfo() string {
	return v.info
}
