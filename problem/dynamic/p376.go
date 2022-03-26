package main

import "fmt"

func main() {
	nums := []int{1, 7, 4, 9, 2, 5}
	fmt.Println(wiggleMaxLength(nums))
}

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	up := 1
	down := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}

	}
	if down > up {
		return down
	} else {
		return up
	}

}
