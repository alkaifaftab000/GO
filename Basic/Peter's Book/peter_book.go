package main

import "fmt"

func main() {
	var tp int16 // total page
	fmt.Scan(&tp)
	if tp > 1000 {
		return
	}

	var pageArr [8]int16 // pages peter read on each day
	for i := range 7 {
		fmt.Scan(&pageArr[i])
	}
	pageArr[7] = 0

	fmt.Println(FinishOnDay(pageArr, tp))
}

func FinishOnDay(pageArr [8]int16, tp int16) int16 {
	var i int16 = 0
	for tp > 0 {
		tp -= pageArr[i]
		i = (i + 1) % 8
	}
	return i
}
