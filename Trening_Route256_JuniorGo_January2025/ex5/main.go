package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Cars struct {
	Start    int
	End      int
	Capacity int
	Load     int
	Index    int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		zakaz, cars := readInput(in)
		result := assignCars(zakaz, cars)

		for _, res := range result {
			fmt.Fprintf(out, "%d ", res)
		}
		fmt.Fprintln(out)
	}
}

func readInput(in *bufio.Reader) ([]int, []Cars) {
	var n int
	fmt.Fscan(in, &n)
	in.ReadString('\n')

	zakaz := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &zakaz[i])
	}

	fmt.Fscan(in, &n)
	in.ReadString('\n')
	cars := make([]Cars, n)

	for j := 0; j < n; j++ {
		var start, end, capacity int
		fmt.Fscanf(in, "%d %d %d\n", &start, &end, &capacity)
		cars[j] = Cars{
			Start:    start,
			End:      end,
			Capacity: capacity,
			Load:     0,
			Index:    j + 1,
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		if cars[i].Start == cars[j].Start {
			return cars[i].Index < cars[j].Index
		}
		return cars[i].Start < cars[j].Start
	})

	return zakaz, cars
}

func assignCars(zakaz []int, cars []Cars) []int {
	result := make([]int, len(zakaz))

	type Order struct {
		Time  int
		Index int
	}
	orders := make([]Order, len(zakaz))
	for i, time := range zakaz {
		orders[i] = Order{Time: time, Index: i}
	}

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].Time < orders[j].Time
	})

	pointer := 0

	for _, car := range cars {
		for pointer < len(orders) && orders[pointer].Time <= car.End && car.Load < car.Capacity {
			if orders[pointer].Time >= car.Start {
				result[orders[pointer].Index] = car.Index
				car.Load++
			} else {
				result[orders[pointer].Index] = -1
			}
			pointer++
		}
	}

	for pointer < len(orders) {
		result[orders[pointer].Index] = -1
		pointer++
	}

	return result
}
