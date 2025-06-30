package main

import "fmt"

func fmtIntro() {
	//Single line comment
	/*	Multi-line comment
		Another line of multi-line comment */

	// Print functions from the fmt package
	fmt.Print("Hello, World!")   // Print prints without a newline
	fmt.Println("Hello, World!") // Println adds a newline at the end

	fmt.Printf("Hello, %s!\n", "World") // Printf allows formatted output
	// different format specifiers can be used
	fmt.Printf("Integer: %d, Float: %.2f, String: %s\n", 42, 3.14, "example") // %d for integers, %.2f for floats, %s for strings
}
