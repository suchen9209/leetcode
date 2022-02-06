package main

import "fmt"

func main() {
	ratings := []int{1,3,4,5,2}
	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {
	l := len(ratings)
	candies := make([]int, l)
	// --- >
	for i := 0;i<l - 1;i++ {
		if ratings[i + 1] > ratings[i]{
			candies[i + 1] = candies[i] + 1
		}
		fmt.Println(candies)
	}

	fmt.Println("---------")
	// < ---
	for i := l-1;i>=1;i--{
		if ratings[i - 1] > ratings[i] && candies[i - 1] <= candies[i]{
			candies[i - 1] = candies[i] + 1
		}

		fmt.Println(candies)
	}

	candiesNum := l
	for _, i2 := range candies {
		candiesNum += i2
	}

	fmt.Println(candies)
	return candiesNum
}
