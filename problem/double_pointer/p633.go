package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(1000))
}

func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	s := 0
	e := int(math.Sqrt(float64(c)))
	for e >= s {
		tmp := e*e + s*s
		if tmp == c {
			return true
		} else if tmp > c {
			e--
		} else {
			s++
		}
	}
	return false
}
