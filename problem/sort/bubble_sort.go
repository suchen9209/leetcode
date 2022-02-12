package main

import "fmt"

func main() {
	nums := []int{6, 41, 8, 3, 5, 6, 10, 1, 6, 34, 123, 56, 5, 3, 54645, 34, 5676457, 34, 6567, 45, 2, 1, 3, 11}
	bubble_sort(nums)
	fmt.Println(nums)
}

func bubble_sort(nums []int) {
	rb := len(nums) - 1
	for rb > 0 {
		i := 0
		for i < rb {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
			i++
		}
		rb--
	}
}
