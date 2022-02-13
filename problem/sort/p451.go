package main

import (
	"fmt"
	"sort"
)

type s451 [][2]int

func (s s451) Len() int {
	return len(s)
}

func (s s451) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s s451) Less(i int, j int) bool {
	return s[i][1] > s[j][1]
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}

func frequencySort(s string) string {
	r := make(map[int]int)
	for _, i2 := range s {
		r[int(byte(i2))]++
	}
	fmt.Println(r)
	var sortArr [][2]int
	for b, i := range r {
		sortArr = append(sortArr, [2]int{b, i})
	}
	sort.Sort(s451(sortArr))
	var rs string
	for _, si := range sortArr {
		for j := 0; j < si[1]; j++ {
			rs += string(byte(si[0]))
		}
	}
	return rs
}
