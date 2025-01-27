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

	numbers := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &numbers[i])
	}
	result := PodschetZp(numbers)

	for _, num := range result {
		fmt.Fprintln(out, num)
	}
}

func PodschetZp(numbers []string) (zp []string) {
	for _, num := range numbers {
		if len(num) < 2 {
			zp = append(zp, "0")
		} else if CheckNumber(num) == 1 {
			zp = append(zp, RemoveFirstChar(num))
		} else if CheckNumber(num) == 2 {
			zp = append(zp, RemoveLastChar(num))
		} else if CheckNumber(num) == 3 {
			zp = append(zp, RemoveMinDigit(num))
		}
	}
	return
}

func CheckNumber(n string) int {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(n); i++ {
		if n[i] > n[i-1] {
			isDecreasing = false
		} else if n[i] < n[i-1] {
			isIncreasing = false
		}
	}

	if isIncreasing {
		return 1
	} else if isDecreasing {
		return 2
	}
	return 3
}

func RemoveFirstChar(n string) string {
	return n[1:]
}

func RemoveLastChar(n string) string {
	return n[:len(n)-1]
}

func RemoveMinDigit(n string) string {
	resultStr := ""
	for i := 0; i < len(n)-1; i++ {
		if n[i] < n[i+1] {
			resultStr = n[:i] + n[i+1:]
			break
		}
	}
	return resultStr
}
