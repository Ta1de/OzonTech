package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	allNumbers := make([]int, t)
	allLines := make([]string, t)
	allCheckLines := make([]string, t)

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &allNumbers[i])
		in.ReadString('\n')

		allLines[i], _ = in.ReadString('\n')

		allCheckLines[i], _ = in.ReadString('\n')
	}

	result := CheckFormat(allNumbers, allLines, allCheckLines)
	for _, num := range result {
		fmt.Fprintln(out, num)
	}
}

func CheckFormat(numbers []int, Line []string, checkLine []string) (zp []string) {
	for i := 0; i < len(numbers); i++ {
		if len(Line[i]) == len(checkLine[i]) {

			if ChekSort(Line[i], checkLine[i]) {
				zp = append(zp, "yes")
			} else {
				zp = append(zp, "no")
			}
		} else {
			zp = append(zp, "no")
		}
	}

	return zp
}

func ChekSort(line, chek string) bool {
	lineNumbers := convertToNumbers(line)
	chekNumbers := convertToNumbers(chek)

	sort.Ints(lineNumbers)

	if len(lineNumbers) != len(chekNumbers) {
		return false
	}
	for i := range lineNumbers {
		if lineNumbers[i] != chekNumbers[i] {
			return false
		}
	}
	return true
}

func convertToNumbers(s string) []int {
	parts := strings.Fields(s)
	numbers := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			numbers[i] = 0
		}
		numbers[i] = num
	}

	return numbers
}
