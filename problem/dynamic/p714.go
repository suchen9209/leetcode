package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fmt.Println(maxProfit714(prices, 2))
}

func maxProfit714(prices []int, fee int) int {
	l := len(prices)
	buy := make([]int, l)
	sell := make([]int, l)

	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	buy[0] = -prices[0]
	sell[0] = 0
	for i := 1; i < l; i++ {
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		sell[i] = max(buy[i-1]+prices[i]-fee, sell[i-1])

	}

	return sell[l-1]
}
