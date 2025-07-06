package main

import "fmt"

func main() {
	encryptedStringLength := 0
	fmt.Scan(&encryptedStringLength)
	if encryptedStringLength <= 0 || encryptedStringLength >= 56 {
		fmt.Print("Invalid Encrypted String Length Error Code -1")
		return
	}

	encryptedString := ""
	fmt.Scan(&encryptedString)
	fmt.Println(DecodeCipher(encryptedString, encryptedStringLength))

}

func DecodeCipher(encryptedString string, encryptedStringLength int) string {
	uniqueString := ""
	fmt.Println("Entered")
	for i := range encryptedStringLength - 1 {
		if encryptedString[i] != encryptedString[i+1] {
			uniqueString += string(encryptedString[i])

		}
	}
	fmt.Println("Exited")
	uniqueString += string(encryptedString[encryptedStringLength-1])
	fmt.Println(uniqueString)
	return ""
}
