package main

import "fmt"

func main() {
	nums := []int{2, 1, 1, 2, 2, 2, 2}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	min := nums[0]
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Println(left, mid, right, min)
		if nums[left] == nums[mid] && nums[right] == nums[mid] && left != right {
			if nums[mid] < min {
				min = nums[mid]
			}
			left++
			right--
		} else if nums[mid] <= nums[right] {
			if nums[mid] < min {
				min = nums[mid]
			}
			right = mid - 1
		} else {
			if nums[left] < min {
				min = nums[left]
			}
			left = mid + 1
		}
	}
	return min
}
