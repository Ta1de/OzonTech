package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	var inputSource *os.File
// 	var err error

// 	inputSource, err = os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("Ошибка открытия файла:", err)
// 		return
// 	}
// 	defer inputSource.Close()

// 	in := bufio.NewReader(inputSource)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	var t int
// 	fmt.Fscan(in, &t)
// 	for i := 0; i < t; i++ {
// 		var n int
// 		fmt.Fscan(in, &n)
// 		strs := make([]string, n)
// 		for j := 0; j < n; j++ {
// 			var s string
// 			fmt.Fscan(in, &s)
// 			strs[j] = s
// 		}
// 		count := countEqualStrings(strs)
// 		fmt.Fprintln(out, count)
// 	}
// }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		strs := make([]string, n)
		for j := 0; j < n; j++ {
			var s string
			fmt.Fscan(in, &s)
			strs = append(strs, s)
		}
		count := countEqualStrings(strs)
		fmt.Fprintln(out, count)
	}
}

func SimilarLines(str1, str2 string) bool {
	// Если длины строк разные, они не могут быть равны
	if len(str1) != len(str2) {
		return false
	}

	// Проверяем, равны ли все четные буквы
	evenEqual := true
	for i := 0; i < len(str1); i++ {
		if i%2 == 0 && str1[i] != str2[i] {
			evenEqual = false
			break
		}
	}

	// Проверяем, равны ли все нечетные буквы
	oddEqual := true
	for i := 0; i < len(str1); i++ {
		if i%2 != 0 && str1[i] != str2[i] {
			oddEqual = false
			break
		}
	}

	// Строки равны, если выполняется хотя бы одно из условий
	return evenEqual || oddEqual
}

func countEqualStrings(arr []string) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if SimilarLines(arr[i], arr[j]) {
				count++
			}
		}
	}
	return count
}
