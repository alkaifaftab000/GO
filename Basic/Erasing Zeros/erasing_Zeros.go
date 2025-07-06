package main

import (
	"fmt"
)

func main() {

	totalTestCases := 0
	var stringsArray []string
	fmt.Scan(&totalTestCases)

	for range totalTestCases {
		a := ""
		fmt.Scan(&a)
		stringsArray = append(stringsArray, a)
	}

	for i := range totalTestCases {
		fmt.Println(ErasingZeros(stringsArray[i]))
	}
}

func ErasingZeros(desiredString string) int {
	length := len(desiredString)
	j := length - 1
	if j == 0 {
		return 0
	}
	zerosBetween1 := 0
	t := 0

	for t < j {
		if t < length && desiredString[t] != '1' {
			t++
		} else if j >= 0 && desiredString[j] != '1' {
			j--
		} else {
			break
		}
	}
	for i := t; i < j; i++ {
		if desiredString[i] == '0' {
			zerosBetween1++
		}
	}
	return zerosBetween1
}
