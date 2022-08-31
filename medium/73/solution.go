package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	solution := setZeroes(matrix)
	fmt.Println(solution)
}

func setZeroes(matrix [][]int) [][]int {

	m := len(matrix)
	n := len(matrix[0])

	row := make(map[int]bool)
	col := make(map[int]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	// Make all rows 0
	for k, _ := range row {
		for i := 0; i < n; i++ {
			matrix[k][i] = 0
		}
	}

	// Make all cols 0
	for k, _ := range col {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}

	return matrix
}
