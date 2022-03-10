package main

import "fmt"

func main() {
	fmt.Println(numSquares(10))
}

func numSquares(n int) int {
	dpMap := make(map[int]int)
	for i := 0; i <= n; i++ {
		dpMap[i] = n
	}
	dpMap[0] = 0

	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dpMap[i] = min(dpMap[i], dpMap[i-j*j]+1)
		}
	}

	return dpMap[n]
}
