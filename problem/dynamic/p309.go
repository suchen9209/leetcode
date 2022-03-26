package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit309(prices))
}

func maxProfit309(prices []int) int {
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	s1 := make([]int, len(prices))
	s2 := make([]int, len(prices))

	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	buy[0] = -prices[0]
	s1[0] = buy[0]
	s2[0] = 0
	sell[0] = 0

	for i := 1; i < len(prices); i++ {
		buy[i] = s2[i-1] - prices[i]
		s1[i] = max(buy[i-1], s1[i-1])
		sell[i] = max(buy[i-1]+prices[i], s1[i-1]+prices[i])
		s2[i] = max(sell[i-1], s2[i-1])
	}

	return max(sell[len(prices)-1], s2[len(prices)-1])
}
