package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	minPrice := 1 << 32
	maxGet := 0
	for _, price := range prices {
		minPrice = min(price, minPrice)
		maxGet = max(maxGet, price-minPrice)
	}
	return maxGet
}
