package main

import "fmt"

func main() {
	var word string
	fmt.Scanf("%s", &word)
	fmt.Print(wordCapital(word))
}

func wordCapital(word string) string {
	if len(word) == 0 {
		return ""
	}
	if word[0] >= 'a' && word[0] <= 'z' {
		return string(word[0]-'a'+'A') + word[1:]
	}
	return word
}
