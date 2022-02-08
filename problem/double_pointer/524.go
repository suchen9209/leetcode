package main

import (
	"fmt"
	"sort"
	"strings"
)

type ss524 []string

func (s ss524) Len() int {
	return len(s)
}

func (s ss524) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ss524) Less(i int, j int) bool {
	if len(s[i]) == len(s[j]) {
		return strings.Compare(s[i], s[j]) < 0
	} else {
		return len(s[i]) > len(s[j])
	}
}

func main() {
	s := "abpcplea"
	dictionary := []string{"ale", "apple", "monkey", "plea"}
	fmt.Println(findLongestWord(s, dictionary))
}

func findLongestWord(s string, dictionary []string) string {
	sort.Sort(ss524(dictionary))
	fmt.Println(dictionary)
	for _, s2 := range dictionary {
		p1, p2 := 0, 0
		for p1 < len(s) && p2 < len(s2) {
			if s[p1] != s2[p2] {
				p1++
			} else {
				if p2 == len(s2)-1 {
					return s2
				}
				p1++
				p2++
			}
		}
	}
	//fmt.Println(strings.Compare("ad","ac"))
	return ""
}
