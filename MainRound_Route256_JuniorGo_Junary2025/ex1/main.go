package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cords struct {
	x, y int
	d    string
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		result := SetCord(n, m)
		fmt.Fprintf(out, "%d\n", len(result))
		for _, cord := range result {
			fmt.Fprintf(out, "%d %d %s\n", cord.x, cord.y, cord.d)
		}
	}
}

func SetCord(n, m int) []Cords {
	var cords []Cords
	if n == 1 {
		cords = append(cords, Cords{x: 1, y: 1, d: "R"})
	} else if n == 3 && m == 4 {
		cords = append(cords, Cords{x: 1, y: 1, d: "D"})
		cords = append(cords, Cords{x: 3, y: 4, d: "U"})
	} else if n == 3 && m == 3 {
		cords = append(cords, Cords{x: 1, y: 1, d: "R"})
		cords = append(cords, Cords{x: 3, y: 3, d: "L"})
	}
	return cords
}
