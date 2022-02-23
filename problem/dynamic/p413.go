package main

import "fmt"

/**
https://leetcode-cn.com/problems/arithmetic-slices/solution/hua-dong-chuang-kou-dong-tai-gui-hua-jav-3vpp/
*/
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(numberOfArithmeticSlices(nums))
}

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	dpMap := make(map[int]int)
	dpMap[0] = 0
	dpMap[1] = 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dpMap[i] = dpMap[i-1] + 1
		}
	}
	sum := 0
	for _, i2 := range dpMap {
		sum += i2
	}
	return sum

}
