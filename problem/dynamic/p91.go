package main

import "fmt"

func main() {
	s := "10"
	fmt.Println(numDecodings(s))
}

func numDecodings(s string) int {

	dpMap := make(map[int]int)
	if s[0] == '0' {
		return 0
	}
	dpMap[0] = 1
	if len(s) <= 1 {
		return 1
	}
	if s[1] == '0' && s[0] > '2' {
		return 0
	}
	if s[0] == '1' && s[1] != '0' {
		dpMap[1] = 2
	} else if s[0] == '2' && s[1] != '0' {
		if s[1] <= '6' {
			dpMap[1] = 2
		} else {
			dpMap[1] = 1
		}
	} else {
		dpMap[1] = 1
	}

	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dpMap[i] = dpMap[i-2]
			} else {
				return 0
			}
		} else {
			if s[i-1] == '1' {
				dpMap[i] = dpMap[i-1] + dpMap[i-2]
			} else if s[i-1] == '2' && s[i] <= '6' {
				dpMap[i] = dpMap[i-1] + dpMap[i-2]
			} else {
				dpMap[i] = dpMap[i-1]
			}
		}
		fmt.Println(dpMap)
	}

	return dpMap[len(s)-1]

}
