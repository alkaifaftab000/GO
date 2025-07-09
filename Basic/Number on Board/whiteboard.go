package main

import (
	"fmt"
	"math"
)

type TestCase struct {
	n        int
	sequence [][]int
}

func main() {
	var t int
	fmt.Scan(&t)

	testCases := make([]TestCase, t)

	// Read all test cases
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		testCases[i] = TestCase{n: n}
	}

	// Process each test case
	for i := range testCases {
		processTestCase(&testCases[i])
	}

	// Output all results
	for _, tc := range testCases {
		fmt.Println(2) // The minimal answer is always 2
		for _, op := range tc.sequence {
			fmt.Println(op[0], op[1])
		}
	}
}

func processTestCase(tc *TestCase) {
	n := tc.n
	sequence := make([]int, n)
	for i := range n {
		sequence[i] = i + 1
	}

	tc.sequence = make([][]int, 0, n-1)

	for len(sequence) > 1 {
		// Always take the two largest numbers
		a := sequence[len(sequence)-1]
		b := sequence[len(sequence)-2]

		// Store the operation
		tc.sequence = append(tc.sequence, []int{a, b})

		// Replace them with their average (ceiling)
		newNum := int(math.Ceil(float64(a+b) / 2))
		sequence = sequence[:len(sequence)-2] // Remove last two
		sequence = append(sequence, newNum)   // Add new number
	}
}
