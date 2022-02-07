package main

import (
	"fmt"
	"sort"
)

type ss406 [][]int

func (s ss406) Len() int {
	return len(s)
}

func (s ss406) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less V3/**
func (s ss406) Less(i, j int) bool {
	if s[i][0] > s[j][0] {
		return true
	} else if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}
	return false
}

func main() {
	s := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(s))
}

/**
V3
*/
func reconstructQueue(people [][]int) [][]int {
	sort.Sort(ss406(people))
	fmt.Println(people)
	var result [][]int
	for _, person := range people {
		if person[1] == 0 {
			result = append([][]int{person}, result[:]...)
		} else {
			fmt.Println(person)
			result = append(result, []int{0})
			copy(result[person[1]+1:], result[person[1]:])
			result[person[1]] = person
			//tmp := result[person[1]:]
			//fmt.Println(tmp)
			//fmt.Println(result[:person[1]])
			//result = append(result[:person[1]],person)
			//fmt.Println(result)
			//result = append(result,tmp...)
			//fmt.Println(result)
			//fmt.Println()
			//tmp := append(result[:person[1]],person)
			//result = append(tmp,result[person[1]:]...)
		}

	}
	return result
}
