package main

import "fmt"

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	dpMap := make(map[[2]int]int)
	if text1[0] == text2[0] {
		dpMap[[2]int{0, 0}] = 1
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dpMap[[2]int{i, j}] = dpMap[[2]int{i - 1, j - 1}] + 1
			} else {
				dpMap[[2]int{i, j}] = max(dpMap[[2]int{i - 1, j}], dpMap[[2]int{i, j - 1}])
			}
		}
	}

	return dpMap[[2]int{len(text1) - 1, len(text2) - 1}]
}
