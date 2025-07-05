package main

import "fmt"

func main() {
	totalCubes := 0
	fmt.Scan(&totalCubes)

	fmt.Print(totalHeight(totalCubes))

}

func totalHeight(totalCubesLeft int) int {
	height := 0
	cubesRequired := 0

	for (cubesRequired + height + 1) <= totalCubesLeft {
		height++
		// fmt.Println("Height: ", height)
		cubesRequired += height
		// fmt.Println("Cubes Required: ", cubesRequired)
		totalCubesLeft -= cubesRequired
		// fmt.Println("Cubes Left: ", totalCubesLeft)
	}

	return height
}
