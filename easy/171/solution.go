package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var col string
	fmt.Print("Enter the column name: ")
	fmt.Scanln(&col)

	count := GetColumnCount(col)
	fmt.Printf("The count for column [%s] is [%v]\n", col, count)
}

func GetColumnCount(col string) int {
	var count float64

	col = strings.ToUpper(col) // Removing this line reduces the memory consumption in-case your input is always caps
	for i, val := range col {
		charCount := val - 64 // 64, because the ascii for A is 65. So, 65-64 is 1, which is the position of A in alphabet
		count += float64(charCount) * math.Pow(float64(26), float64(len(col)-(i+1)))
	}
	return int(count)
}
