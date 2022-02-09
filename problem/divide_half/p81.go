package main

import "fmt"

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}

func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	a := -1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			a = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return a != -1 && nums[a] == target
}

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[mid] <= nums[right] {
			if binarySearch(nums[mid:right+1], target) {
				return true
			}
			right = mid - 1
		} else {
			if binarySearch(nums[left:mid+1], target) {
				return true
			}
			left = mid + 1
		}

	}
	return false
}
