package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	maxLen := 0
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	dpMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		dpMap[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dpMap[i] = max(dpMap[j]+1, dpMap[i])
			}
		}
	}
	for _, i2 := range dpMap {
		maxLen = max(maxLen, i2)
	}

	return maxLen
}
