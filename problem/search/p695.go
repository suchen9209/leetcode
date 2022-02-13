package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid))

}

func maxAreaOfIsland(grid [][]int) int {
	had := make(map[[2]int]int)
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if had[[2]int{i, j}] != 1 && grid[i][j] != 0 {
				r := findArea(grid, i, j, had)
				if r > maxArea {
					maxArea = r
				}
			}
		}
	}

	fmt.Println(had)
	return maxArea
}

func findArea(grid [][]int, i int, j int, had map[[2]int]int) int {
	if grid[i][j] == 0 {
		return 0
	} else {
		had[[2]int{i, j}] = 1
		area := 1
		//左
		if j-1 >= 0 && had[[2]int{i, j - 1}] != 1 {
			area += findArea(grid, i, j-1, had)
		}
		//下
		if i+1 < len(grid) && had[[2]int{i + 1, j}] != 1 {
			area += findArea(grid, i+1, j, had)
		}
		//右
		if j+1 < len(grid[i]) && had[[2]int{i, j + 1}] != 1 {
			area += findArea(grid, i, j+1, had)
		}
		//上
		if i-1 >= 0 && had[[2]int{i - 1, j}] != 1 {
			area += findArea(grid, i-1, j, had)
		}
		return area
	}

}
