package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println(findMaxForm(strs, m, n))
}

func findMaxForm(strs []string, m int, n int) int {
	lens := len(strs)
	dp := make([][][]int, lens+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}

	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Println(lens)
	for i := 1; i <= lens; i++ {
		zeros := strings.Count(strs[i-1], "0")
		ones := len(strs[i-1]) - zeros
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if j >= zeros && k >= ones {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeros][k-ones]+1)
				} else {
					fmt.Println(i, j, k)
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[lens][m][n]
}
