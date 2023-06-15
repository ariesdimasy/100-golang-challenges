package main

import "fmt"

func multiDimensionSlice(x int, y int) [][]int {
	var result = [][]int{}

	for i := 0; i < x; i++ {
		var row = []int{}
		for j := 0; j < y; j++ {
			row = append(row, i*j)
		}
		result = append(result, row)
	}

	return result
}

func main() {
	fmt.Println(multiDimensionSlice(3, 5))
}
