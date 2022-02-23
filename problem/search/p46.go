package main

import "fmt"

func main() {
	nums := []int{0, 1}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	var result [][]int
	var rs func([]int, int)
	rs = func(nums []int, k int) {
		if k == len(nums)-1 {
			nums3 := make([]int, len(nums))
			copy(nums3, nums)
			result = append(result, nums3)
			return
		}

		for i := k; i < len(nums); i++ {
			nums[i], nums[k] = nums[k], nums[i]
			rs(nums, k+1)
			nums[i], nums[k] = nums[k], nums[i]
		}
	}

	rs(nums, 0)
	return result
}
