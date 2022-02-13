package main

import "fmt"

func main() {
	heights := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	fmt.Println(pacificAtlantic(heights))
}

func pacificAtlantic(heights [][]int) [][]int {

	var result [][]int
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			fmt.Println(i, j)
			visited := make(map[[2]int]int)
			visited2 := make(map[[2]int]int)
			if canGo(heights, i, j, 1, visited) && canGo(heights, i, j, 2, visited2) {
				result = append(result, []int{i, j})
			}
		}

	}
	return result
}

func canGo(heights [][]int, i int, j int, t int, visited map[[2]int]int) bool {
	//t: 1 - 太;2 - 西
	if (i == 0 || j == 0) && t == 1 {
		return true
	}
	if (i == len(heights)-1 || j == len(heights[0])-1) && t == 2 {
		return true
	}
	visited[[2]int{i, j}] = 1
	r := false
	//fmt.Println(visited)
	//fmt.Println(i,j,t)
	//left
	if j-1 >= 0 && visited[[2]int{i, j - 1}] != 1 && heights[i][j-1] <= heights[i][j] {
		r = r || canGo(heights, i, j-1, t, visited)
	}
	//right
	if j+1 < len(heights[0]) && visited[[2]int{i, j + 1}] != 1 && heights[i][j+1] <= heights[i][j] {
		r = r || canGo(heights, i, j+1, t, visited)
	}
	//up
	if i-1 >= 0 && visited[[2]int{i - 1, j}] != 1 && heights[i-1][j] <= heights[i][j] {
		r = r || canGo(heights, i-1, j, t, visited)
	}
	//down
	if i+1 < len(heights) && visited[[2]int{i + 1, j}] != 1 && heights[i+1][j] <= heights[i][j] {
		r = r || canGo(heights, i+1, j, t, visited)
	}
	return r
}
