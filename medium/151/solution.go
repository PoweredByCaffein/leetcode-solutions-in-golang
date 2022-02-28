package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the input string: ")
	scanner.Scan()
	input = scanner.Text()
	fmt.Printf("Input : %s\n", input)

	output := reverseWords(input)
	fmt.Printf("Output : [%s]\n", output)
}

func reverseWords(s string) string {
	// trim whitespaces
	s = strings.TrimSpace(s)

	strSlice := strings.Split(s, " ")
	var result string
	for i := len(strSlice) - 1; i >= 0; i-- {
		if strSlice[i] == "" || strSlice[i] == " " {
			continue
		}
		fmt.Println(strSlice[i])
		result += strSlice[i] + " "
	}
	result = strings.TrimRight(result, " ")
	return result
}
