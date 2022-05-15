package main

import (
	"fmt"
	"math"
)

func main() {
	strs := []string{"ab", "a"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {
	var letter, result string
	var maxIteration int
	maxIteration = math.MaxInt32

	for x := 0; x < maxIteration; x++ {
		for k, v := range strs {

			if maxIteration > len(v) {
				maxIteration = len(v)
				fmt.Println(maxIteration)
			}

			if v == "" {
				return ""
			}

			if k == 0 {
				letter = v[x : x+1]
			}

			if letter != v[x:x+1] {
				return result
			}

		}

		result += letter

		if x == len(strs[0])-1 {
			return result
		}

	}
	return result
}
