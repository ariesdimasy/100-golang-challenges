package main

import "fmt"

func computeFactorial(num int) int {
	var res = 1
	for i := 2; i <= num; i++ {
		res = res * i
	}
	return res
}

func main() {
	fmt.Println(computeFactorial(8))
	fmt.Println(computeFactorial(3))
}
