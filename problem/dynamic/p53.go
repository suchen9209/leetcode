package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxInt := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxInt = max(dp[i], maxInt)
	}

	return maxInt

}
