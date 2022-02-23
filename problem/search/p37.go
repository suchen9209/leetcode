package main

import "fmt"

func main() {
	fmt.Println(byte(0))
	fmt.Println('1' - '1')
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}

func solveSudoku2(board [][]byte) {
	var checkCanset func(int, int, byte) bool
	checkCanset = func(i int, j int, i3 byte) bool {
		for h := 0; h < 9; h++ {

			if i3 == board[i][h] {
				return false
			}
			if i3 == board[h][j] {
				return false
			}
		}
		areaX := i / 3
		areaY := j / 3
		for k := areaX * 3; k < areaX*3+3; k++ {
			for k2 := areaY * 3; k2 < areaY*3+3; k2++ {
				//fmt.Println(board[k][k2],i3,i3 == board[k][k2])
				if i3 == board[k][k2] {
					return false
				}
			}
		}
		return true
	}

	var dfs func(int, int)

	endBoard := make([][]byte, 9)
	dfs = func(i int, j int) {
		fmt.Println(i, j)
		if i > 8 {
			return
		}
		if j > 8 {
			dfs(i+1, 0)
			return
		}
		if i == 8 && j == 8 {
			copy(endBoard, board)
		}
		if board[i][j] == '.' {
			for num := 1; num <= 9; num++ {
				if checkCanset(i, j, byte(num)) {
					board[i][j] = byte(num)
					dfs(i, j+1)
					board[i][j] = '.'
				}
			}
		} else {
			dfs(i, j+1)

		}
	}
	dfs(0, 0)
	fmt.Println(endBoard)
}
