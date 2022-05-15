package main

import "fmt"

func main() {
	x := 121

	fmt.Println(isPalindromeNumber(x))

}

func isPalindromeNumber(n int) bool {

	reverse := getReverseNumber(n)
	if n == reverse {
		return true
	}
	return false
}

func getReverseNumber(n int) int {

	var reverse int
	for n > 0 {
		reverse = (reverse * 10) + (n % 10)
		n = n / 10
	}

	return reverse

}
