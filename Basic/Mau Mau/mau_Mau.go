package main

import "fmt"

func main() {
	var tableCard string
	var cards [5]string

	fmt.Scanln(&tableCard)

	for i := range 5 {
		fmt.Scan(&cards[i])
	}
	fmt.Print(CardMatch(tableCard, cards[:]))
}

func CardMatch(tableCard string, cards []string) string {
	for i := range 5 {
		if (cards[i][0] == tableCard[0]) || (cards[i][1] == tableCard[1]) {
			return "YES"
		}
	}
	return "NO"
}
