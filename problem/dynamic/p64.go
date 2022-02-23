package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	min := func(i int, j int) int {
		if i < j {
			return i
		} else {
			return j
		}
	}

	dpMap := make(map[[2]int]int)
	dpMap[[2]int{0, 0}] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j > 0 {
				dpMap[[2]int{i, j}] = dpMap[[2]int{i, j - 1}] + grid[i][j]
			}
			if j == 0 && i > 0 {
				dpMap[[2]int{i, j}] = dpMap[[2]int{i - 1, j}] + grid[i][j]
			}
			if i > 0 && j > 0 {
				dpMap[[2]int{i, j}] = min(dpMap[[2]int{i - 1, j}]+grid[i][j], dpMap[[2]int{i, j - 1}]+grid[i][j])
			}

		}
	}
	return dpMap[[2]int{len(grid) - 1, len(grid[0]) - 1}]
}
