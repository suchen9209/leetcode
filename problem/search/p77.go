package main

import "fmt"

func main() {
	//fmt.Println(combine2(1,1))
	fmt.Println(combine(1, 1))
}

func combine(n int, k int) (result [][]int) {
	var nums []int
	var ff func(int)
	ff = func(i int) {
		if len(nums)+(n-i+1) < k {
			return
		}
		if len(nums) == k {
			tmp := make([]int, k)
			copy(tmp, nums)
			result = append(result, tmp)
			return
		}
		nums = append(nums, i)
		ff(i + 1)
		nums = nums[:len(nums)-1]
		ff(i + 1)
	}
	ff(1)
	return
}

var result77 [][]int
var nums77 []int

func combine2(n int, k int) [][]int {

	recall(1, n, k)
	return result77
}

func recall(min int, n int, k int) {
	fmt.Println(min, nums77)
	if (len(nums77) + (n - min + 1)) < k {
		return
	}
	if len(nums77) == k {
		tmp := make([]int, k)
		copy(tmp, nums77)
		result77 = append(result77, tmp)
		return
	}

	nums77 = append(nums77, min)
	recall(min+1, n, k)
	nums77 = nums77[:len(nums77)-1]
	recall(min+1, n, k)
}
