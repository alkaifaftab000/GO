package main

import "fmt"

func main() {
	inputString := ""
	fmt.Scan(&inputString)
	inputStringLength := len(inputString)
	fmt.Println(RemoveVowels(UpperToLower(inputString, inputStringLength), inputStringLength))

}

func UpperToLower(inputString string, inputStringLength int) string {
	lowerCaseString := []rune(inputString)
	for i := range inputStringLength {
		if lowerCaseString[i] >= 'A' && lowerCaseString[i] <= 'Z' {
			lowerCaseString[i] = lowerCaseString[i] - 'A' + 'a'
		}
	}
	return string(lowerCaseString)
}

func RemoveVowels(inputString string, inputStringLength int) string {
	var resultRune []rune
	consonentString := []rune(inputString)
	for i := range inputStringLength {
		switch consonentString[i] {
		case 'a', 'e', 'i', 'o', 'u', 'y':
			continue
		default:
			resultRune = append(resultRune, '.', consonentString[i])
		}
	}
	return string(resultRune)
}
