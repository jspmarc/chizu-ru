package parse

import (
	"bufio"
	"bytes"
	"chizu-ru/graph"
	"chizu-ru/node"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fnameBuffer := new(bytes.Buffer)
	fmt.Fprintf(fnameBuffer, "../input/%s", input)
	file, ferr := os.Open(fnameBuffer.String())

	if ferr != nil {
		panic(ferr)
	}

	scanner1 := bufio.NewScanner(file)
	f := make([]string, 0)
	g := graph.New()
	for scanner1.Scan() {
		line := scanner1.Text()
		f = append(f, line)
	}

	nVertex, _ := strconv.Atoi(f[0])
	ctr := 0
	fctr := 1
	for ctr < nVertex {
		splitted := strings.Split(f[fctr], " ")
		firstcoor, _ := strconv.ParseFloat(splitted[1], 64)
		secondcoor, _ := strconv.ParseFloat(splitted[2], 64)
		node.New(splitted[0], firstcoor, secondcoor)
		ctr++
		fctr++
	}
}
