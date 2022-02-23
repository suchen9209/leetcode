package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	fmt.Println(combinationSum2(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	fmt.Println(candidates)

	var tranString func([]int) string
	tranString = func(ints []int) (s string) {
		for _, i2 := range ints {
			s += strconv.Itoa(i2)
			s += "-"
		}
		return s
	}

	var dfs func(int)
	var sum int
	var numArr []int
	var haveMap map[string]int
	haveMap = make(map[string]int)
	dfs = func(i int) {
		if sum > target {
			return
		} else if sum == target {
			if haveMap[tranString(numArr)] != 1 {
				tmpArr := make([]int, len(numArr))
				copy(tmpArr, numArr)
				result = append(result, tmpArr)
				haveMap[tranString(numArr)] = 1
			}
			return
		} else {
			if i == len(candidates) {
				return
			}
			for j := i; j < len(candidates); j++ {
				sum += candidates[j]

				numArr = append(numArr, candidates[j])
				fmt.Println(numArr)
				dfs(j + 1)
				sum -= candidates[j]
				numArr = numArr[:len(numArr)-1]
			}
		}
	}
	dfs(0)

	return result
}
