package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "SEE"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	var wordArr []byte
	for _, w := range word {
		wordArr = append(wordArr, byte(w))
	}
	r := false
	var nums []byte
	var dfs func(int, int)
	visited := make(map[[2]int]int)
	dfs = func(i int, j int) {
		if string(nums) == string(wordArr) {
			r = true
			return
		}
		if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || len(nums) >= len(wordArr) || visited[[2]int{i, j}] == 1 {
			return
		}
		if board[i][j] == wordArr[len(nums)] {
			nums = append(nums, board[i][j])
			visited[[2]int{i, j}] = 1
			fmt.Println(string(nums), string(wordArr))

			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
			nums = nums[:len(nums)-1]
			visited[[2]int{i, j}] = 0
		} else {
			return
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j)
		}
	}

	return r
}
