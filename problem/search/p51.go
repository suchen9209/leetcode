package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(5))
}

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}

	var dfs func(int)
	var canSetQueen func(int, int) bool
	var setQueenMap = make(map[[2]int]int)
	var nowStringArr []string
	var result [][]string
	var returnQueenString func(int) string
	returnQueenString = func(k int) string {
		tmps := make([]byte, n)
		for j := 0; j < n; j++ {
			tmps[j] = '.'
		}
		tmps[k] = 'Q'
		return string(tmps)
	}
	canSetQueen = func(i int, j int) bool {
		key := [2]int{i, j}
		if setQueenMap[key] == 1 {
			return false
		}
		for h := 1; h <= n; h++ {
			if setQueenMap[[2]int{h, j}] == 1 || setQueenMap[[2]int{i, h}] == 1 {
				return false
			}
			if i-h >= 1 && j-h >= 1 && setQueenMap[[2]int{i - h, j - h}] == 1 {
				return false
			}
			if i+h <= n && j-h >= 1 && setQueenMap[[2]int{i + h, j - h}] == 1 {
				return false
			}
			if i-h >= 1 && j+h <= n && setQueenMap[[2]int{i - h, j + h}] == 1 {
				return false
			}
			if i+h <= n && j+h <= n && setQueenMap[[2]int{i + h, j + h}] == 1 {
				return false
			}
		}
		return true
	}

	dfs = func(row int) {

		if row == n {
			for j := 1; j <= n; j++ {
				if canSetQueen(row, j) {
					nowStringArr = append(nowStringArr, returnQueenString(j-1))
					tmp := make([]string, len(nowStringArr))
					copy(tmp, nowStringArr)
					result = append(result, tmp)
					nowStringArr = nowStringArr[:len(nowStringArr)-1]
					return
				}
			}
			return
		}

		for k := 1; k <= n; k++ {
			if canSetQueen(row, k) {
				setQueenMap[[2]int{row, k}] = 1
				nowStringArr = append(nowStringArr, returnQueenString(k-1))
				dfs(row + 1)
				nowStringArr = nowStringArr[:len(nowStringArr)-1]
				setQueenMap[[2]int{row, k}] = 0
			}
		}
		return

	}

	dfs(1)

	return result
}
