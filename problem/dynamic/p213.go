package main

import "fmt"

func main() {
	nums := []int{2, 3, 3, 3, 23, 23, 23, 3}
	fmt.Println(rob2(nums))
}

func rob2(nums []int) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else if len(nums) == 3 {
		return max(max(nums[0], nums[1]), nums[2])
	}
	nums1 := nums[:len(nums)-1]
	nums2 := nums[1:]
	fmt.Println(nums1, nums2)
	dp1 := make([]int, len(nums1))
	dp2 := make([]int, len(nums2))
	dp1[0] = nums1[0]
	dp1[1] = max(nums1[0], nums1[1])
	dp2[0] = nums2[0]
	dp2[1] = max(nums2[0], nums2[1])
	for i := 2; i < len(nums1); i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums1[i])
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums2[i])
	}

	return max(dp1[len(nums1)-1], dp2[len(nums2)-1])
}
