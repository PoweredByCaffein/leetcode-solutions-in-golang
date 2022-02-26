package main

//import "fmt"
//
//func main() {
//	var input int
//	fmt.Print("Enter the number you want to find the trailing zeroes for: ")
//	fmt.Scanln(&input)
//
//	output := trailingZeroes(input)
//	fmt.Printf("Trailing zeroes for [%v] is [%v]\n", input, output)
//}
//
//func trailingZeroes(input int) int {
//
//	factorial := getFactorial(input)
//	fmt.Println("Factorial", getFactorial(input))
//
//	numTwos := getDigitCount(factorial, 2)
//	numFives := getDigitCount(factorial, 5)
//
//	if numTwos == 0 || numFives == 0 {
//		return 0
//	} else {
//		if numTwos < numFives {
//			return numTwos
//		}
//		return numFives
//	}
//}
//
//func getFactorial(input int) int {
//	factorial := input
//
//	for {
//		if input == 1 {
//			return factorial
//		}
//		factorial = factorial * (input - 1)
//		input--
//	}
//}
//
//func getDigitCount(input, digit int) int {
//	var digitCount int
//	for {
//		if input%digit == 0 {
//			input = input / digit
//			digitCount++
//		} else {
//			return digitCount
//		}
//	}
//}
