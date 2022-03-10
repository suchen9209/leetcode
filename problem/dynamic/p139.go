package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}

	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	dpMap := make(map[int]bool)
	dpMap[0] = true
	for i := 1; i <= len(s); i++ {
		for _, s2 := range wordDict {
			if i >= len(s2) {
				fmt.Println(s[i-len(s2):i], s2)
			}

			if i >= len(s2) && s[i-len(s2):i] == s2 {
				dpMap[i] = dpMap[i] || dpMap[i-len(s2)]
			}
		}
	}

	return dpMap[len(s)]
}
