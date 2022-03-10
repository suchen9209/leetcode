package main

import "fmt"

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	n := len(word1)
	m := len(word2)
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[0][j] = j
		}
		dp[i][0] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}

	return dp[n][m]
}

func minDistance2(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1) //dp数组存放上一次dp的情况
	for i := 1; i < len(word2)+1; i++ {
		dp[i] = i //①:对第1行进行初始化,此时取word1前0个字母与word2前i个字母的编辑距离 即word2的长度
	}
	for i := 1; i < len(word1)+1; i++ {
		leftUp := i - 1 //用于存储dp[i][j]左上点的数据(因为会被dp[i-1][j]覆盖掉)
		dp[0] = i       //初始化每一行的首位元素 类似①的情况
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] { //当此处字符相同 则通过替换的转移方式可以实现无替换转移,即不变
				dp[j], leftUp = min(dp[j]+1, dp[j-1]+1, leftUp), dp[j]
			} else {
				dp[j], leftUp = min(dp[j]+1, dp[j-1]+1, leftUp+1), dp[j]
			}
		}
	}
	return dp[len(word2)]
}

func min(a, b, c int) (ret int) {
	ret = a
	if a > b {
		ret = b
		if b > c {
			ret = c
		}
	} else if a > c {
		ret = c
	}
	return
}
