package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	a, b, c := 0, 0, 0
	for i := 0; i < n; i++ {
		input(&a, &b, &c)
		fmt.Println(exhandling(a, b, c))
	}
}

func input(a *int, b *int, c *int) {
	fmt.Scan(a, b, c)
}
func exhandling(a, b, c int) int {
	i := 0
	for a <= c && b <= c {
		if a > b {
			b += a
		} else {
			a += b
		}
		i++
	}
	return i
}
