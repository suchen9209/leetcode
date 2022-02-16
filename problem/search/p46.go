package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

var result [][]int

func permute(nums []int) [][]int {
	rs(nums, 0)
	return result
}

func rs(nums []int, k int) {
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
