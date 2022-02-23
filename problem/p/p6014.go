package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "aababab"
	repeatLimit := 2
	fmt.Println(repeatLimitedString(s, repeatLimit))
}

func repeatLimitedString(s string, repeatLimit int) string {
	charMap := make(map[int]int)
	var charSortArr []int
	for _, i2 := range s {
		charMap[int(i2)]++
	}
	for i := range charMap {
		charSortArr = append(charSortArr, i)
	}
	sort.Ints(charSortArr)
	fmt.Println(charSortArr)

	var findSmall func(int) int
	findSmall = func(i int) int {
		for j := i - 1; j >= 0; j-- {
			if charMap[charSortArr[j]] > 0 {
				charMap[charSortArr[j]]--
				return charSortArr[j]
			}
		}
		return -1
	}

	var nowNum int
	var nowCount int
	returnString := ""
	for i := len(charSortArr) - 1; i >= 0; i-- {
		nowCount = 0
		nowNum = charSortArr[i]
		fmt.Println(nowNum)
		fmt.Println(nowCount)
		for j := 0; j < charMap[charSortArr[i]]; j++ {
			if nowNum != charSortArr[i] {
				nowCount = 1
			}
			nowNum = charSortArr[i]
			if nowCount == repeatLimit {
				tmp := findSmall(i)
				if tmp < 0 {
					goto Loop
				} else {
					returnString += string(byte(tmp))
					nowNum = tmp
					nowCount = 1
				}
			}
			returnString += string(byte(charSortArr[i]))
			nowCount++
		}
	}
Loop:
	return returnString
}

func repeatLimitedString2(s string, repeatLimit int) string {
	var arr []int
	for _, i2 := range s {
		arr = append(arr, int(i2))
	}
	sort.Ints(arr)
	var findSmall func(int, int) int
	findSmall = func(i int, s int) int {
		for j := i; j >= 0; j-- {
			if arr[j] < s {
				return j
			}
		}
		return -1
	}

	var needDelete []int
	var nowNum int
	var nowCount int
	for i := len(arr) - 1; i >= 0; i-- {
		if nowNum == arr[i] {
			if nowCount >= repeatLimit {
				findPos := findSmall(i, arr[i])
				if findPos < 0 {
					needDelete = append(needDelete, i)
				} else {
					arr[i], arr[findPos] = arr[findPos], arr[i]
					nowNum = arr[i]
					nowCount = 1
				}
			} else {
				nowCount++
			}
		} else {
			nowNum = arr[i]
			nowCount = 1
		}
	}

	for _, i2 := range needDelete {
		var tmp []int
		tmp = make([]int, len(arr)-1)
		tmp = arr[:i2]
		tmp = append(tmp, arr[i2+1:]...)
		arr = tmp
	}

	returnString := ""
	for i := len(arr) - 1; i >= 0; i-- {
		returnString += string(byte(arr[i]))
	}

	return returnString
}
