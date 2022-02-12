package main

import "fmt"

func main() {
	nums := []int{6, 41, 8, 3, 5, 6, 10, 1, 6, 34, 123, 56, 5, 3, 54645, 34, 5676457, 34, 6567, 45, 2, 1, 3, 11}
	quick_sort(nums)
	fmt.Println(nums)
}

func quick_sort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	mid := nums[0]
	i := 1
	left, right := 0, len(nums)-1
	for left < right {
		if nums[i] > mid {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		}
	}
	quick_sort(nums[:left])
	quick_sort(nums[left+1:])

}
