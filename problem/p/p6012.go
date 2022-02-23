package main

import "fmt"

func main() {
	fmt.Println(countEven(1000))
}

func countEven(num int) int {
	var count int
	if num == 1000 {
		num = 999
	}
	for i := 1; i <= num; i++ {
		qn := i / 100
		qn2 := i % 100
		qn3 := qn2 / 10
		qn4 := qn2 % 10

		if (qn+qn3+qn4)%2 == 0 {
			count++
		}

	}
	return count
}
