package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	bucket := [3]int{}
	for _, num := range nums {
		bucket[num]++
	}
	k := 0
	for i, i2 := range bucket {
		for j := 0; j < i2; j++ {
			nums[k] = i
			k++
		}
	}
}
