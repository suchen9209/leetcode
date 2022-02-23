package main

import "fmt"

func main() {
	fmt.Println(climbStairs(50))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var m map[int]int
	m = make(map[int]int)
	m[1] = 1
	m[2] = 2
	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n]
}
