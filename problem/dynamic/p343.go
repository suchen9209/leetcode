package main

import "fmt"

func main() {
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 2; i <= n; i++ {
		imax := 0
		for j := 1; j < i; j++ {
			imax = max(imax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = imax
	}
	return dp[n]
}
