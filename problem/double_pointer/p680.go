package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		fmt.Println(string(s[left]), string(s[right]))
		if s[left] == s[right] {
			left++
			right--
		} else {
			l1, r1, is1 := left, right-1, true
			for l1 < r1 {
				if s[l1] == s[r1] {
					l1++
					r1--
				} else {
					is1 = false
					break
				}
			}
			l2, r2, is2 := left+1, right, true
			for l2 < r2 {
				if s[l2] == s[r2] {
					l2++
					r2--
				} else {
					is2 = false
					break
				}
			}
			return is1 || is2
		}

	}
	return true
}
