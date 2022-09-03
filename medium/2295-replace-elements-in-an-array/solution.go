package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 6}
	operations := [][]int{
		{1, 3},
		{4, 7},
		{6, 1},
	}

	fmt.Println(arrayChange(nums, operations))

}

func arrayChange(nums []int, operations [][]int) []int {
	// Create an arrayMap to store all the elements with location
	// By doing so we don't have to search for the index that needs to be updated
	arrayMap := make(map[int]int)
	for k, v := range nums {
		arrayMap[v] = k
	}

	// Update the nums array
	for _, operation := range operations {
		// Store the value of index so that we can update our map
		temp := arrayMap[operation[0]]

		// Update the nums array
		nums[arrayMap[operation[0]]] = operation[1]

		// Update the map with the latest value
		arrayMap[operation[1]] = temp
	}

	return nums
}
