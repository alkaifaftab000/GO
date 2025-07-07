package main

import "fmt"

func main() {
	totalTestCase := 0
	fmt.Scan(&totalTestCase)
	var sharedAnswerArray []int

	for i := 0; i < totalTestCase; i++ {
		sharedAnswerArray = append(sharedAnswerArray, 0)
	}

	for i := 0; i < totalTestCase; i++ {
		rows, cols, grid := input() // Get grid for this test case
		solve(rows, cols, grid, sharedAnswerArray, i)
	}

	for i := 0; i < totalTestCase; i++ {
		fmt.Println(sharedAnswerArray[i])
	}
}

func input() (int, int, [][]byte) {
	var rows, cols int
	fmt.Scan(&rows, &cols)

	grid := make([][]byte, rows)
	for i := range grid {
		grid[i] = make([]byte, cols)
	}

	for i := 0; i < rows; i++ {
		var ch string
		fmt.Scan(&ch)
		for j := 0; j < cols; j++ {
			if ch[j] != 'R' && ch[j] != 'C' && ch[j] != 'D' {
				return -1, -1, nil
			}
			grid[i][j] = ch[j]
		}
	}
	return rows, cols, grid
}

func solve(rows, cols int, grid [][]byte, ans []int, testCaseIdx int) {
	count := 0
	for j := 0; j < cols-1; j++ {
		if grid[rows-1][j] == 'D' {
			count++
		}
	}
	for i := 0; i < rows-1; i++ {
		if grid[i][cols-1] == 'R' {
			count++
		}
	}
	ans[testCaseIdx] = count
}
