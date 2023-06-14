package main

import (
	"fmt"
	"strconv"
)

func findNumberDivisible() string {
	var result = ""
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			result = result + strconv.Itoa(i) + ","
		}
	}

	return result
}

func main() {
	fmt.Println(findNumberDivisible())
}
