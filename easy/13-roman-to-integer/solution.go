package main

import "fmt"

func main() {
	roman := "XVX"
	fmt.Println(romanToInt(roman))
}

func romanToInt(s string) int {
	roman := make(map[string]int)
	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000

	var result int
	for i := 0; i < len(s); {
		if (i < len(s)-1) && (roman[string(s[i+1])] > roman[string(s[i])]) {
			result += (roman[string(s[i+1])] - roman[string(s[i])])
			i += 2
		} else {
			result += roman[string(s[i])]
			i += 1
		}
	}
	return result
}
