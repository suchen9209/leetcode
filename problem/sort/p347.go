package main

import (
	"fmt"
	"sort"
)

type s347 [][]int

func (s s347) Len() int {
	return len(s)
}

func (s s347) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less V3/**
func (s s347) Less(i, j int) bool {
	return s[i][1] >= s[j][1]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent2(nums, 2))
}

func topKFrequent2(nums []int, k int) []int {

	total := make(map[int]int)
	for _, num := range nums {
		total[num]++
	}
	fmt.Println(total)
	var keys [][]int
	for i, i2 := range total {
		keys = append(keys, []int{i, i2})
	}
	fmt.Println(keys)
	sort.Sort(s347(keys))
	fmt.Println(keys)
	var final []int
	i := 0
	for i < k {
		final = append(final, keys[i][0])
		i++
	}
	return final
}

func topKFrequent(nums []int, k int) []int {

	var total []int
	for _, num := range nums {
		if num > len(total)-1 {
			for i := len(total); i <= num+1; i++ {
				total = append(total, 0)
			}
		}
		fmt.Println(total)
		total[num]++
	}
	fmt.Println(total)
	var final []int
	for len(final) < k {
		maxKey, i := 0, 0
		for i < len(total) {
			if total[i] > total[maxKey] {
				maxKey = i
			}
			i++
		}
		final = append(final, maxKey)
		total[maxKey] = 0
	}
	return final
}
