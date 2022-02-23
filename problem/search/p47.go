package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	var perm []int
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(k int) {
		if k == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			dfs(k + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	dfs(0)
	return
}

func permuteUnique2(nums []int) [][]int {
	var result [][]int
	var rs func(int)
	rs = func(k int) {
		if k == len(nums)-1 {
			nums3 := make([]int, len(nums))
			copy(nums3, nums)
			result = append(result, nums3)
			return
		}

		for i := k; i < len(nums); i++ {
			nums[i], nums[k] = nums[k], nums[i]
			rs(k + 1)
			nums[i], nums[k] = nums[k], nums[i]
		}
	}
	rs(0)

	return result
}
