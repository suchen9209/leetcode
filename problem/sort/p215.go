package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	now := 0
	for now < len(nums) {
		maxKey := now
		i := now
		for i < len(nums) {
			if nums[i] > nums[maxKey] {
				maxKey = i
			}
			i++
		}
		nums[maxKey], nums[now] = nums[now], nums[maxKey]
		fmt.Println(nums)
		now++
		if now == k {
			return nums[now-1]
		}

	}
	return nums[now-1]
}
