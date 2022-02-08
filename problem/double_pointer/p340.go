package main

import "fmt"

//TODO leetcode未测试
func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1))
}

func lengthOfLongestSubstringKDistinct(s string, t int) string {
	left, right := 0, 0
	nowDiffNum := 0
	charDict := map[byte]int{}

	charDict[s[left]]++
	nowDiffNum++

	var stringArr []string

	for right < len(s) {
		if nowDiffNum > t {
			charDict[s[left]]--
			if charDict[s[left]] == 0 {
				nowDiffNum--
			}
			left++
		} else {
			if nowDiffNum == t {
				stringArr = append(stringArr, s[left:right+1])
			}
			right++
			if right < len(s) {
				charDict[s[right]]++
				if charDict[s[right]] == 1 {
					nowDiffNum++
				}
			}
		}
	}

	maxString := ""
	for _, s3 := range stringArr {
		if len(s3) > len(maxString) {
			maxString = s3
		}
	}

	return maxString
}
