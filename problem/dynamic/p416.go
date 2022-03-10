package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 == 1 {
		return false
	}
	half := sum / 2
	if max > half {
		return false
	}

	n := len(nums)

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, half+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true

	for i := 1; i < n; i++ {
		for j := 1; j <= half; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][half]
}

func canPartition2(nums []int) bool {
	sum := 0
	max := 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 == 1 {
		return false
	}
	half := sum / 2
	if max > half {
		return false
	}
	dpMap := make(map[[2]int]bool)
	for i := 0; i < len(nums); i++ {
		dpMap[[2]int{i, 0}] = true
	}
	dpMap[[2]int{0, nums[0]}] = true

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= half; j++ {
			if j >= nums[i] {
				dpMap[[2]int{i, j}] = dpMap[[2]int{i - 1, j}] || dpMap[[2]int{i - 1, j - nums[i]}]
			} else {
				dpMap[[2]int{i, j}] = dpMap[[2]int{i - 1, j}]
			}

		}
	}
	for i := 0; i <= len(nums); i++ {
		for j := 0; j <= half; j++ {
			fmt.Println(i+1, j, dpMap[[2]int{i, j}])

		}
	}

	return dpMap[[2]int{len(nums) - 1, half}]
}
