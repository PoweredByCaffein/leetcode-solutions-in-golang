package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter the number you want to find the trailing zeroes for: ")
	fmt.Scanln(&input)

	output := trailingZeroes(input)
	fmt.Printf("Trailing zeroes for [%v] is [%v]\n", input, output)
}

func trailingZeroes(input int) int {

	var numTwos, numFives int

	for i := 1; i <= input; i++ {
		numTwos += getDigitCount(i, 2)
		numFives += getDigitCount(i, 5)
	}

	if numTwos < numFives {
		return numTwos
	}
	return numFives

}

func getDigitCount(input, digit int) int {
	var digitCount int
	for {
		if input%digit == 0 {
			input = input / digit
			digitCount++
		} else {
			return digitCount
		}
	}
}
