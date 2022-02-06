package main

import (
	"fmt"
	"sort"
)
type ss452 [][]int

func (s ss452) Len() int {
	return len(s)
}

func (s ss452) Swap(i int,j int){
	s[i],s[j] = s[j],s[i]
}

// Less V3/**
func (s ss452) Less(i,j int) bool {
	if s[i][1] < s[j][1]{
		return true
	}else if s[i][1] == s[j][1]{
		return s[i][0] < s[j][0]
	}
	return false
}

func main() {
	//s := [][]int{{-52,31},{-73,-26},{82,97},{-65,-11},{-62,-49},{95,99},{58,95},{-31,49},{66,98},{-63,2},{30,47},{-40,-26}}
	s := [][]int{{1,2},{3,4},{5,6},{7,8}}
	fmt.Println(findMinArrowShots(s))
}

/**
V3
*/
func findMinArrowShots(intervals [][]int) int {
	sort.Sort(ss452(intervals))
	remove := 0
	nowKey := 0
	for i := 1; i < len(intervals); i++ {
		fmt.Println(intervals[i],intervals[nowKey])
		fmt.Println(nowKey)
		if intervals[i][0] <= intervals[nowKey][1]{
			remove ++
		}else{
			nowKey = i
		}
	}

	return len(intervals) - remove
}
