package main

import "fmt"

func main() {
	nums := []int{12, 6, 41, 8, 3, 5, 6, 10, 1, 6, 34, 123, 56, 5, 3, 54645, 34, 5676457, 34, 6567, 45, 2, 1, 3, 11}
	select_sort(nums)
	fmt.Println(nums)
}

func select_sort(nums []int) {
	now := 0
	for now < len(nums)-1 {
		minKey := now
		i := now + 1
		for i < len(nums) {
			if nums[i] < nums[minKey] {
				minKey = i
			}
			i++
		}
		nums[now], nums[minKey] = nums[minKey], nums[now]
		now++
		fmt.Println(nums)
	}

}
