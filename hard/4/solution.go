package main

import (
	"fmt"
)

func main() {
	nums1 := []int{0, 0}
	nums2 := []int{0, 0}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	finalSlice := append(nums1, nums2...)

	fmt.Println(finalSlice)
	if len(finalSlice)%2 == 0 {
		x := float64(finalSlice[len(finalSlice)/2-1])
		y := float64(finalSlice[len(finalSlice)/2])

		return (x + y) / 2
	} else {
		x := float64(finalSlice[len(finalSlice)/2])
		return x
	}

}
