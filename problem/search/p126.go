package main

import "fmt"

func main() {
	beginWord := "hot"
	endWord := "dog"
	wordList := []string{"hot", "dog", "dot"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var result [][]string
	var checkTransOne func(string, string) bool
	checkTransOne = func(s string, s2 string) bool {
		diff := 0
		for i := range s {
			if s[i] != s2[i] {
				diff++
			}
		}
		return diff == 1
	}
	have := false
	for _, s := range wordList {
		if endWord == s {
			have = true
			break
		}
	}
	if !have {
		return result
	}
	if len(beginWord) == 1 {
		result = append(result, []string{beginWord, endWord})
		return result
	}

	var queue []string
	var nowString string
	var selected map[string]int
	selected = make(map[string]int)
	nowString = beginWord
	for _, s := range wordList {
		if checkTransOne(nowString, s) {
			queue = append(queue, s)
			selected[s] = 1
		}
	}
	fmt.Println(queue)
	step := 1
	canFind := false
	for len(queue) > 0 {
		len := len(queue)
		for i := 0; i < len; i++ {
			s := queue[i]
			for _, s2 := range wordList {
				if s2 == beginWord {
					continue
				}
				if selected[s2] != 1 && checkTransOne(s, s2) {
					selected[s] = 1
					queue = append(queue, s2)
				}
				if s == endWord {
					canFind = true
					goto Loop
				}
			}
		}
		step++
		queue = queue[len:]
		fmt.Println(queue)
	}
Loop:
	fmt.Println(step)
	if canFind {
		var road []string
		var findNext func(string, int)
		road = append(road, beginWord)
		findNext = func(s string, i int) {
			if i == 0 {
				return
			}
			if i == 1 {
				for _, s2 := range wordList {
					if checkTransOne(s, s2) && endWord == s2 {
						road = append(road, s2)
						tmpArr := make([]string, len(road))
						copy(tmpArr, road)
						result = append(result, tmpArr)
						road = road[:len(road)-1]
						return
					}
				}
				return
			}
			for _, s2 := range wordList {
				if checkTransOne(s, s2) {
					road = append(road, s2)
					findNext(s2, i-1)
					road = road[:len(road)-1]
				}
			}
		}
		findNext(beginWord, step)

	}

	return result

}
