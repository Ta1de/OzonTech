package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		fmt.Fprintln(out, CountTablets(n))
	}
}

func CountTablets(num int) int {
	res := 0
	if num <= 9 {
		res = num + 1
	} else if num == 10 {
		res = 10
	} else {
		res = Tables(num)
	}
	return res
}

func Tables(num int) int {
	count := (len(strconv.Itoa(num)) - 1) * 10
	numstr := strconv.Itoa(num)
	dopnum := int(numstr[0] - '0')

	srstr := ""
	for i := 0; i < len(numstr); i++ {
		srstr += strconv.Itoa(dopnum)
	}
	srnum, _ := strconv.Atoi(srstr)

	if num >= srnum {
		count += dopnum
	} else {
		count += dopnum - 1
	}

	return count
}
