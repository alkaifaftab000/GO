package main // Package main is the entry point of the program
import "fmt" // fmt is the package for formatted I/O

func main() {
	var a int = 0
	var b int = 0
	fmt.Scanf("%d %d", &a, &b)        // Read two integers from standard input
	fmt.Println(BearAndbrother(a, b)) // Call the BearAndbrother function with example arguments
}

func BearAndbrother(num1 int, num2 int) int {
	count := 0
	for num1 <= num2 {
		num1 = num1 * 3
		num2 = num2 * 2
		count++
	}
	return count
}
