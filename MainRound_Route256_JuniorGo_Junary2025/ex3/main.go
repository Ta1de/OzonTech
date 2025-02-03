package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		m1 := make(map[string]string)
		for j := 0; j < n; j++ {
			var a, b string
			fmt.Fscan(in, &a, &b)
			m1[a] = b
		}
		var s string
		fmt.Fscan(in, &s)
		data := ParseString(s)
		if Compare(m1, data) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func Compare(m map[string]string, data []string) bool {
	priceToItems := make(map[string][]string)
	for item, price := range m {
		priceToItems[price] = append(priceToItems[price], item)
	}

	seenPrices := make(map[string]bool)
	seenItems := make(map[string]bool)

	for i := 0; i < len(data); i += 2 {
		if i+1 >= len(data) {
			return false
		}

		item, price := data[i], data[i+1]

		itemsWithPrice, exists := priceToItems[price]
		if !exists {
			return false
		}

		found := false
		for _, validItem := range itemsWithPrice {
			if validItem == item {
				found = true
				break
			}
		}
		if !found {
			return false
		}

		if seenPrices[price] {
			return false
		}
		seenPrices[price] = true

		if seenItems[item] {
			return false
		}
		seenItems[item] = true
	}

	for price, items := range priceToItems {
		if !seenPrices[price] {
			if len(items) > 1 {
				return false
			}
		}
	}

	for _, items := range priceToItems {
		if len(items) == 1 {
			if _, exists := seenItems[items[0]]; !exists {
				return false
			}
		}
	}

	return true
}

func ParseString(s string) []string {
	pars := strings.Split(s, ",")
	var data []string
	for _, v := range pars {
		kv := strings.Split(v, ":")
		data = append(data, kv...)
	}
	return data
}
