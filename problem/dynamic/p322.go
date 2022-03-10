package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = amount + 1
		}
		dp[i][0] = 0
	}
	dp[0][0] = amount + 1

	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			//
			//dp[i][j] = min(dp[i][j-coins[i-1]] + 1,dp[i-1][j])
			//dp[j] = min(dp[j-coins[i-1]]+1,dp[j])
			//if j >= coins[i-1]{
			//	dp[i][j] = min(dp[i][j-coins[i-1]]+1,dp[i-1][j])
			//}else{
			//	dp[i][j] = dp[i-1][j]
			//}

			if j >= coins[i-1] {
				dp[i][j] = min(dp[i][j-coins[i-1]]+1, dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[len(coins)][amount] > amount {
		return -1
	} else {
		return dp[len(coins)][amount]
	}

}
