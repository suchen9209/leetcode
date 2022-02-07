package main

import "fmt"

func main() {
	nums := []int{5, 7, 1, 8}
	fmt.Println(checkPossibility(nums))
}

func checkPossibility(nums []int) bool {
	LessNum := 0
	removeKey := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			removeKey = i
			LessNum++
		}
		if LessNum >= 2 {
			return false
		}
	}
	fmt.Println(removeKey)

	returnBool1 := true
	returnBool2 := true
	for i := 0; i < len(nums)-1; i++ {
		if removeKey == i {
			if i-1 >= 0 && nums[i+1] < nums[i-1] {
				returnBool1 = false
			}
		} else if removeKey == i-1 {
			if i-1 >= 0 && i+1 < len(nums) && nums[i+1] < nums[i-1] {
				returnBool2 = false
			}
		} else {
			if nums[i+1] < nums[i] {
				returnBool1 = false
			}
		}

	}

	fmt.Println(returnBool2)
	return returnBool1 || returnBool2
}
