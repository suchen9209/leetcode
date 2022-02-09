package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRangeV3(nums, target))
}

/**
V3
*/
func searchRangeV3(nums []int, target int) []int {
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
	if a == -1 || nums[a] != target {
		return []int{-1, -1}
	}

	left2, right2 := 0, len(nums)-1
	b := -1
	for left2 <= right2 {
		mid := left2 + (right2-left2)/2
		if nums[mid] < target {
			left2 = mid + 1
		} else {
			b = mid
			right2 = mid - 1
		}
	}

	return []int{b, a}

}

/**
官方题解
*/
func searchRangeV2(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

/**
V1
*/
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	a := -1
	r := []int{-1, -1}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			a = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println(a)
	fmt.Println(nums[a])
	if a == -1 || nums[a] != target {
		return r
	}

	for i := a; i >= 0; i-- {
		if nums[i] != target {
			r[0] = i + 1
			break
		} else {
			r[0] = i
		}
	}
	for j := a; j < len(nums); j++ {
		if nums[j] != target {
			r[1] = j - 1
			break
		} else {
			r[1] = j
		}
	}

	return r
}
