package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	fmt.Println(findContentChildren(g, s))
}

/**

 */
func findContentChildrenV2(greed, size []int) (ans int) {

	sort.Ints(greed)
	sort.Ints(size)
	n, m := len(greed), len(size)
	for i, j := 0, 0; i < n && j < m; i++ {
		for j < m && greed[i] > size[j] {
			j++
		}
		if j < m {
			ans++
			j++
		}
	}
	return
}

/**
v1
*/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	num := 0
	for _, i1 := range g {
		if len(s) == 0 {
			return num
		}
		for k, i2 := range s {
			if i2 >= i1 {
				num++
				if len(s) == 1 {
					return num
				}
				s = append(s[:k], s[k+1:]...)
				break
			}
			if k == len(s)-1 {
				return num
			}
		}

	}

	return num
}
