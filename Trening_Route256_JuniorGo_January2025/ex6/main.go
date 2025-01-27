package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	A, B int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		// pairs := readInput(in)
		// for _, res := range result {
		// 	fmt.Fprintf(out, "%d ", res)
		// }
		// fmt.Fprintln(out)
	}
}

// func readInput(in *bufio.Reader) []Pair {
// 	var n, m int
// 	fmt.Fscan(in, &n, &m)
// 	in.ReadString('\n')

// 	pairs := make([]Pair, 0, m)

// 	for i := 0; i < m; i++ {
// 		var a, b int
// 		fmt.Fscan(in, &a, &b)
// 		pairs = append(pairs, Pair{A: a, B: b})
// 	}
// 	return pairs
// }

// func Checklever() {

// }
