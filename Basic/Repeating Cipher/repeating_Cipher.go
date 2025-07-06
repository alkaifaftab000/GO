package main

import "fmt"

func main() {
	encryptedStringLength := 0
	fmt.Scan(&encryptedStringLength)
	if encryptedStringLength <= 0 || encryptedStringLength >= 56 {
		return
	}
	encryptedString := ""
	fmt.Scan(&encryptedString)
	fmt.Println(DecodeCipher(encryptedString, encryptedStringLength))

}

func DecodeCipher(encryptedString string, encryptedStringLength int) string {
	uniqueString := ""
	sum := 0
	i := 0
	for sum < encryptedStringLength {
		uniqueString += string(encryptedString[sum])
		i++
		sum += i
	}
	return uniqueString
}
