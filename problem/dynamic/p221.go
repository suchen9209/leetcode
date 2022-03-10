package main

import "fmt"

func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	dp := make(map[[2]int]int)
	if matrix[0][0] == '1' {
		dp[[2]int{0, 0}] = 1
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
	maxLen := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[[2]int{i, j}] = min(min(dp[[2]int{i - 1, j}], dp[[2]int{i, j - 1}]), dp[[2]int{i - 1, j - 1}]) + 1
			}
			maxLen = max(maxLen, dp[[2]int{i, j}])

		}
	}
	return maxLen * maxLen
}
