package main

import (
	"fmt"
	"sort"
)

func main() {
	pairs := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(findLongestChain(pairs))
}

type ss646 [][]int

func (s ss646) Len() int {
	return len(s)
}

func (s ss646) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ss646) Less(i int, j int) bool {
	if s[i][1] < s[j][1] {
		return true
	} else {
		return false
	}
}

func findLongestChain(pairs [][]int) int {
	sort.Sort(ss646(pairs))
	fmt.Println(pairs)
	maxLen := 0
	for i := 0; i < len(pairs); i++ {
		nowMaxEnd := pairs[i][1]
		nowMaxLen := 1
		for j := i + 1; j < len(pairs); j++ {
			if pairs[j][0] > nowMaxEnd {
				nowMaxLen++
				nowMaxEnd = pairs[j][1]
			}
		}
		if maxLen < nowMaxLen {
			maxLen = nowMaxLen
		}
	}
	return maxLen
}
