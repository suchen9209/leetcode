package main

import "fmt"

//[1 1 2 3 3 3 5 5 6 6 6 8 10 11 34 34 34 41 45 56 123 6567 54645 5676457]
func main() {
	nums := []int{6, 41, 8, 3, 5, 6, 10, 1, 6, 34, 123, 56, 5, 3, 54645, 34, 5676457, 34, 6567, 45, 2, 1, 3, 11}
	fmt.Println(merge_sort(nums))
}

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := merge_sort(nums[:mid])
	right := merge_sort(nums[mid:])
	return merge(left, right)
}

func merge(nums1, nums2 []int) (result []int) {
	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] <= nums2[p2] {
			result = append(result, nums1[p1])
			p1++
		} else {
			result = append(result, nums2[p2])
			p2++
		}
	}
	result = append(result, nums1[p1:]...)
	result = append(result, nums2[p2:]...)
	return result
}
