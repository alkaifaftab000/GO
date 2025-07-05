package main

import "fmt"

func main() {
	totalTestCases := 0
	var appartmentNumbers []int
	fmt.Scan(&totalTestCases)
	for range totalTestCases {
		a := 0
		fmt.Scan(&a)
		appartmentNumbers = append(appartmentNumbers, a)
	}
	for i := range totalTestCases {
		fmt.Println(DigitsPressed(appartmentNumbers[i]))
	}
}

func DigitsPressed(residentAppartmentNumber int) int {
	if residentAppartmentNumber == 1 {
		return 1
	}
	totalDigitsPressed := 0
	for i := 1; i <= (residentAppartmentNumber % 10); i++ {
		power := 1
		j := i

		if i == (residentAppartmentNumber % 10) {
			for j <= residentAppartmentNumber {
				fmt.Println("Power: ", power)
				switch power {
				case 1:
					totalDigitsPressed++
				case 10:
					totalDigitsPressed += 2
				case 100:
					totalDigitsPressed += 3
				case 1000:
					totalDigitsPressed += 4
				}
				fmt.Println("Total Digits Pressed: ", totalDigitsPressed)
				power *= 10
				j = j + (i * power)
			}
		} else {
			for j < 10000 {
				fmt.Println("Power: ", power)
				switch power {
				case 1:
					totalDigitsPressed++
				case 10:
					totalDigitsPressed += 2
				case 100:
					totalDigitsPressed += 3
				case 1000:
					totalDigitsPressed += 4
				}
				fmt.Println("Total Digits Pressed: ", totalDigitsPressed)
				power *= 10
				j = j + (i * power)
			}
		}
	}
	return totalDigitsPressed
}
