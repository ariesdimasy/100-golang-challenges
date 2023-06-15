package main

import (
	"fmt"
	"strings"
)

func sortingName(str string) []string {
	var ssplit = strings.Split(str, ",")
	for i := 0; i < len(ssplit); i++ {
		var temp1 = ""
		for j := i; j < len(ssplit); j++ {
			if ssplit[j][0] < ssplit[i][0] {
				temp1 = ssplit[i]
				ssplit[i] = ssplit[j]
				ssplit[j] = temp1
			}
		}
	}

	return ssplit
}

func sortingNameSelection(str string) []string {
	var ssplit = strings.Split(str, ",")
	for i := 0; i < len(ssplit); i++ {
		var min = ssplit[i]
		var positionMin = i
		var temp = ""
		for j := i; j < len(ssplit); j++ {
			if ssplit[j] < min {
				min = ssplit[j]
				positionMin = j
			}
		}

		if min < ssplit[i] {
			temp = ssplit[i]
			ssplit[i] = min
			ssplit[positionMin] = temp
		}
	}

	return ssplit
}

func main() {
	fmt.Println(sortingName("bag,hello,without,world"))
	fmt.Println(sortingNameSelection("bag,hello,without,world"))
	//var name string = "dimas"
	//fmt.Println(name[1] > name[0])
}
