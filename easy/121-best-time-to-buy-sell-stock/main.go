package main

import "fmt"

func main() {
	prices := []int{1, 2, 7, 3, 4, 10}
	fmt.Println("Input:", prices)
	fmt.Println("Max Profit:", maxProfit(prices))
}

func maxProfit(prices []int) int {
	profit, minSP := 0, prices[0]
	for _, price := range prices {
		if minSP > price {
			minSP = price
		}

		temp := price - minSP
		if temp > profit {
			profit = temp
		}
	}
	return profit
}
