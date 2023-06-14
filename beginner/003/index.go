package main

import "fmt"

func createMap(n int) map[int]int {
	var result = map[int]int{}

	for i := 1; i <= n; i++ {
		result[i] = i * i
	}

	return result
}

func main() {
	fmt.Println(createMap(4))
}
