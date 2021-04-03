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
	"bytes"
	//"chizu-ru/graph"
	"chizu-ru/node"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertex = node.Node

// Parse fungsi untuk parsing file menjadi graf kemudian pointer graf akan
// dikembalikan
func Parse() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fnameBuffer := new(bytes.Buffer)
	fmt.Fprintf(fnameBuffer, "../test/%s", input)
	file, ferr := os.Open(fnameBuffer.String())

	if ferr != nil {
		panic(ferr)
	}

	scanner1 := bufio.NewScanner(file)
	f := make([]string, 0)
	for scanner1.Scan() {
		line := scanner1.Text()
		f = append(f, line)
	}
	defer file.Close()

	// buat nyimpen sudut yang udah dibuat, jadi tinggal tambah ke graf
	vertices := make([]*vertex, 0)

	nVertex, _ := strconv.Atoi(f[0])

	for ctr := 0; ctr < nVertex; ctr++ {
		splitted := strings.Split(f[ctr+1], " ")
		latitude, _ := strconv.ParseFloat(splitted[1], 64)
		longitude, _ := strconv.ParseFloat(splitted[2], 64)

		vertices = append(vertices, node.New(splitted[0], latitude, longitude))
	}

	fmt.Println(vertices)
	for _, v := range vertices {
		fmt.Println(v)
	}
}