package main

import "fmt"

func main() {
	//board := [][]byte{{'X','X','X','X'},{'X','O','O','X'}, {'X','X','O','X'},{'X','O','X','X'}}
	board := [][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}
	solve(board)
	fmt.Println(board)
}

func solve(board [][]byte) {
	var visitBoard map[[2]int]int
	visitBoard = make(map[[2]int]int)
	var goNext func(int, int)
	goNext = func(i int, j int) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
			return
		}
		if board[i][j] == 'O' && visitBoard[[2]int{i, j}] != 1 {
			visitBoard[[2]int{i, j}] = 1
			goNext(i+1, j)
			goNext(i-1, j)
			goNext(i, j+1)
			goNext(i, j-1)
		} else {
			return
		}
	}
	for i := 0; i < len(board); i++ {

		goNext(i, 0)
		goNext(i, len(board[0])-1)

	}
	for i := 0; i < len(board[0]); i++ {
		goNext(0, i)
		goNext(len(board)-1, i)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' && visitBoard[[2]int{i, j}] != 1 {
				board[i][j] = 'X'
			}
		}
	}
}
