package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	TmpCommission, Commission := 0, 0
	TmpPrice := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] > TmpPrice {
			TmpCommission += prices[i] - TmpPrice
		} else if prices[i] < TmpPrice {
			Commission += TmpCommission
			TmpCommission = 0
		}
		TmpPrice = prices[i]
	}
	return Commission + TmpCommission
}
