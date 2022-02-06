package main

import "fmt"

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	//r := flowerbed[:]
	r := append(flowerbed, 0)
	r = append(r, 1)
	r = append([]int{0, 1}, r...)
	fmt.Println(r)

	left, canSet := 0, 0

	for i := 1; i < len(r); i++ {
		if r[i] == 1 {
			tmpCanSetNumber := (i - left - 2) / 2
			if tmpCanSetNumber > 0 {
				canSet += tmpCanSetNumber
			}
			left = i
		}
	}

	return n <= canSet
}
