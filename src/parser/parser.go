// Package parser untuk parsing file menjadi graf
// format file:
// baris pertama adalah banyak sudut, misalnya N sudut
// baris kedua sampai N + 1 akan berisi informasi sudut (nama, latitude, longitude)
// baris N + 2 sampai akhir akan berisi matriks berukuran N x N yang berisi
// matriks ketetanggaan antarsudut
// misalnya, jika ada 2 sudut Dago dan ITB dengan jarak dan koordinat arbitrary:
// ```
// 2
// Dago 0 0
// ITB 69 420
// 0 10
// 10 10
// ```
package parser

import (
	"bufio"
	"chizu-ru/graph"
	"chizu-ru/node"
	"io"
	"os"
	"strconv"
	"strings"
)

type vertex = node.Node

// Parse fungsi untuk parsing file menjadi graf kemudian pointer graf akan
// dikembalikan
func Parse(fileContent string) (*graph.Graph, error) {
	f := strings.Split(fileContent, "\n")

	// buat nyimpen sudut yang udah dibuat, jadi tinggal tambah ke graf
	nVertex, _ := strconv.Atoi(f[0])
	vertices := make([]*vertex, nVertex)

	ctr := 1
	for ctr <= nVertex {
		splitted := strings.Split(f[ctr], ",")
		latitude, _ := strconv.ParseFloat(splitted[1], 64)
		longitude, _ := strconv.ParseFloat(splitted[2], 64)

		vertices[ctr-1] = node.New(splitted[0], latitude, longitude)
		ctr++
	}

	g := graph.New()
	for i := 0; i < nVertex; i++ {
		splitted := strings.Split(f[ctr], " ")
		for j := i; j < nVertex; j++ {
			if splitted[j] == "0" {
				continue
			} else {
				g.AddEdge(vertices[i], vertices[j])
			}
		}
		ctr++
	}

	return g, nil
}

func ParseFile(pathToFile string) (*graph.Graph, error) {
	file, ferr := os.Open(pathToFile)
	if ferr != nil {
		return nil, ferr
	}

	scanner1 := bufio.NewScanner(file)
	f := make([]string, 0)
	for scanner1.Scan() {
		line := scanner1.Text()
		f = append(f, line)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return Parse(string(content))
}
