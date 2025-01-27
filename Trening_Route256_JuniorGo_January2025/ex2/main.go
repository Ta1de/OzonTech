package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	process := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &process[i])
	}
	result := PodschetZp(process)

	for _, num := range result {
		fmt.Fprintln(out, num)
	}
}

func PodschetZp(process []string) (res []string) {
	for _, pr := range process {
		if pr[0] != 'M' {
			res = append(res, "NO")
		} else if pr[len(pr)-1] != 'D' {
			res = append(res, "NO")
		} else {
			res = append(res, OtherVariant(pr))
		}
	}
	return
}

func OtherVariant(prstr string) string {
	for i := 0; i < len(prstr)-1; i++ {
		curent := prstr[i]
		next := prstr[i+1]

		switch curent {
		case 'M':
			if next != 'R' && next != 'C' && next != 'D' {
				return "NO"
			}
		case 'R':
			if next != 'D' && next != 'C' {
				return "NO"
			}
		case 'C':
			if next != 'M' {
				return "NO"
			}
		case 'D':
			if next != 'M' {
				return "NO"
			}
		}
	}
	return "YES"
}
