package main

import "fmt"

func main() {
	nums := []int{6, 41, 8, 3, 5, 6, 10, 1, 6, 34, 123, 56, 5, 3, 54645, 34, 5676457, 34, 6567, 45, 2, 1, 3, 11}
	insert_sort(nums)
	fmt.Println(nums)
}

func insert_sort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := 1
	for i < len(nums) {
		c := i
		for c > 0 && nums[c-1] > nums[c] {
			nums[c-1], nums[c] = nums[c], nums[c-1]
			c--
		}
		i++
	}
}
