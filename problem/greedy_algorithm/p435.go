package main

import (
	"fmt"
	"sort"
)

type ss [][]int

func (s ss) Len() int {
	return len(s)
}

func (s ss) Swap(i int,j int){
	s[i],s[j] = s[j],s[i]
}

// Less V3/**
func (s ss) Less(i,j int) bool {
	if s[i][1] < s[j][1]{
		return true
	}else if s[i][1] == s[j][1]{
		return s[i][0] < s[j][0]
	}
	return false
}
// Less V2/**
//func (s ss) Less(i,j int) bool {
//	if s[i][0] < s[j][0]{
//		return true
//	}else if s[i][0] == s[j][0]{
//		return s[i][1] < s[j][1]
//	}
//	return false
//}


func main() {
	s := [][]int{{-52,31},{-73,-26},{82,97},{-65,-11},{-62,-49},{95,99},{58,95},{-31,49},{66,98},{-63,2},{30,47},{-40,-26}}
	fmt.Println(eraseOverlapIntervals(s))
}

/**
V3
 */
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(ss(intervals))
	remove := 0
	fmt.Println(intervals)
	nowKey := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[nowKey][1]{
			remove ++
			fmt.Println(intervals[i])
			fmt.Println(nowKey)
		}else{
			nowKey = i
		}
	}

	return remove
}


/**
v2
 */
func eraseOverlapIntervalsV2(intervals [][]int) int {
	sort.Sort(ss(intervals))
	remove := 0
	fmt.Println(intervals)
	nowKey := len(intervals) - 1
	for i := len(intervals) - 2; i > 0; i-- {
		fmt.Println(intervals[i],intervals[nowKey])
		if intervals[i][1] > intervals[nowKey][0]{
			remove ++
			//fmt.Println(intervals[i])
		}else{
			nowKey = i
		}
	}

	return remove
}

/**
v1
 */
func eraseOverlapIntervalsV1(intervals [][]int) int {
	sort.Sort(ss(intervals))
	remove := 0
	fmt.Println(intervals)
	nowKey := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[nowKey][1]{
			remove ++
			fmt.Println(intervals[i])
			fmt.Println(nowKey)
		}else{
			nowKey = i
		}
	}

	return remove
}
