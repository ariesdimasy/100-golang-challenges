package main

import (
	"fmt"
	"strconv"
	"strings"
)

func createSlice(sequence string) []int {
	var sequenceSplit = strings.Split(sequence, ", ")
	var newSlice = []int{}

	for i := 0; i < len(sequenceSplit); i++ {
		sqSplit, err := strconv.Atoi(sequenceSplit[i])
		if err != nil {
			panic("must be string")
		}
		newSlice = append(newSlice, sqSplit)
	}

	return newSlice
}

func main() {
	fmt.Println(createSlice("34, 67, 55, 33, 12, 98"))
}
