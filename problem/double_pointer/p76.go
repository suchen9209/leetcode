package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	length := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < length {
				length = r - l + 1
				ansL, ansR = l, l+length
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

/**
V1
*/
func minWindowV1(s string, t string) string {
	tMap := make(map[int32]int)
	sMap := make(map[int32]int)
	for _, i2 := range t {
		tMap[i2]++
	}
	left, right := 0, 0
	sArr := []int32(s)

	checkHaveChar := func() bool {
		for i, i2 := range tMap {
			if sMap[i] < i2 {
				return false
			}
		}
		return true
	}

	shortString := ""

	for left <= len(s) && right <= len(s) {
		if checkHaveChar() {
			tmpString := string(sArr[left:right])
			fmt.Println(tmpString)
			if len(tmpString) < len(shortString) || shortString == "" {
				shortString = tmpString
			}
			if _, ok := tMap[sArr[left]]; ok {
				sMap[sArr[left]]--
			}
			left++
		} else {
			if right >= len(s) {
				break
			}
			if _, ok := tMap[sArr[right]]; ok {
				sMap[sArr[right]]++
			}
			right++
		}

	}

	return shortString
}
