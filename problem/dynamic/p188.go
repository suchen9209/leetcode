package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit2(2, prices))
}

func maxProfit2(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

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

	k = min(k, n/2)
	buy := make([][]int, n)
	sell := make([][]int, n)
	for i := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}

	maxP := 0
	for _, i2 := range sell[n-1] {
		maxP = max(i2, maxP)
	}
	return maxP
}
