package main

import "fmt"

func main() {
	fmt.Println(minSteps(10))
}

func minSteps(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 1; j < i; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]
}
