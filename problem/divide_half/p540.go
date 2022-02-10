package main

import "fmt"

func main() {
	//nums := []int{1,1,2,3,3,4,4,8,8}
	nums := []int{1, 1, 3, 7, 7, 11, 11}
	fmt.Println(singleNonDuplicate(nums))
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if right-left <= 2 {
			if nums[left] != nums[mid] {
				return nums[left]
			} else {
				return nums[right]
			}
		}
		if nums[mid] == nums[mid-1] {
			if (mid-left+1)%2 == 0 {
				left = mid + 1
			} else {
				right = mid
			}
		} else if nums[mid] == nums[mid+1] {
			if (right-mid+1)%2 == 0 {
				right = mid - 1
			} else {
				left = mid
			}
		} else {
			return nums[mid]
		}
	}
	return 0
}
