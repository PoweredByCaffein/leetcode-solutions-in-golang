package main

import "fmt"

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
}

func generate(numRows int) [][]int {
	var result [][]int
	result = append(result, []int{1})
	for i := 2; i <= numRows; i++ {
		row := []int{1}
		requiredNums := i - 2

		if requiredNums > 0 {
			previousRow := result[i-2]
			for k, v := range previousRow {
				if k != len(previousRow)-1 {
					row = append(row, v+previousRow[k+1])
				}
			}
		}

		row = append(row, 1)
		result = append(result, row)

	}

	return result
}
