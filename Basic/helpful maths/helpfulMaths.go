package main

import "fmt"

func main() {
	var expression string
	count1, count2, count3 := 0, 0, 0
	fmt.Scan(&expression)

	countNumbers(removePlusSigns(expression), &count1, &count2, &count3)
	fmt.Println(reconstructExpression(count1, count2, count3))
}

func removePlusSigns(expression string) string {
	var removedExpression string = ""
	for i := 0; i < len(expression); i++ {
		if expression[i] != '+' {
			removedExpression += string(expression[i])
		}
	}
	return removedExpression
}

func countNumbers(removedPlusSigns string, count1 *int, count2 *int, count3 *int) int {
	for i := 0; i < len(removedPlusSigns); i++ {
		switch removedPlusSigns[i] {
		case '1':
			*count1++
		case '2':
			*count2++
		case '3':
			*count3++
		}
	}
	return *count1 + *count2*2 + *count3*3
}

func reconstructExpression(count1, count2, count3 int) string {
	var finalexpression string
	for range count1 {
		finalexpression += "1+"
	}
	for range count2 {
		finalexpression += "2+"
	}
	for range count3 {
		finalexpression += "3+"
	}
	lengthfe := len(finalexpression)
	if lengthfe > 0 {
		finalexpression = finalexpression[:lengthfe-1]
	}

	return finalexpression
}
