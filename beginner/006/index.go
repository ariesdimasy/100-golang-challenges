package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func calculated(nums string) []float64 {
	var numsSplit = strings.Split(nums, ",")
	var res = []float64{}

	for i := 0; i < len(numsSplit); i++ {

		nSplit, err := strconv.Atoi(numsSplit[i])
		//fmt.Println(" ===> ", nSplit)

		if err != nil {
			panic("must be number")
		}
		floatNum := float64((2 * 50 * nSplit) / 30)
		var q = math.Round(math.Sqrt(floatNum))
		res = append(res, q)
	}

	return res
}

func main() {
	fmt.Println(calculated("1,2,3,4"))
	fmt.Println(calculated("100,150,180"))
}
