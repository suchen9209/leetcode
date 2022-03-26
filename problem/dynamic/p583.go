package main

import "fmt"

func main() {
	word1 := "leetcode"
	word2 := "etco"

	fmt.Println(minDistance583(word1, word2))
}

func minDistance583(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	for i := 0; i <= l2; i++ {
		dp[0][i] = i
	}

	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[l1][l2]
}
