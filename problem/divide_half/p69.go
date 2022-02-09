package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
}

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	l, r := 0, x
	lastMid := -1
	for l < r {
		mid := (l + r) / 2
		product := mid * mid
		if product == x {
			l = mid
			break
		} else if product < x {
			l = mid
		} else {
			r = mid
		}
		if lastMid == mid {
			break
		} else {
			lastMid = mid
		}
	}
	return l
}

/**
官方题解
*/
func mySqrt2(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
