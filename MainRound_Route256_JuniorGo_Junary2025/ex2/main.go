package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		var n int
		fmt.Fscan(in, &n)
		fmt.Fprintln(out, minDigitPlates(n))
	}
}

func minDigitPlates(n int) int {
	count := make([]int, 10)

	for i := n; i >= 0; i-- {
		for _, digit := range strconv.Itoa(i) {
			count[digit-'0']++
		}
	}

	maxPlates := 0
	for _, v := range count {
		if v > maxPlates {
			maxPlates = v
		}
	}

	return maxPlates
}
