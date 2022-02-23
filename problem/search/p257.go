package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var t = new(TreeNode)
	t.Val = 1
	t.Left = new(TreeNode)
	t.Left.Val = 2
	t.Left.Right = new(TreeNode)
	t.Left.Right.Val = 3
	t.Right = new(TreeNode)
	t.Right.Val = 3

	fmt.Println(binaryTreePaths(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var result []string
	var numArr []int
	var resultArr [][]int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			numArr = append(numArr, node.Val)
			var tmpArr = make([]int, len(numArr))
			copy(tmpArr, numArr)
			resultArr = append(resultArr, tmpArr)
			numArr = numArr[:len(numArr)-1]
			return
		}
		if node.Left != nil {
			numArr = append(numArr, node.Val)
			dfs(node.Left)
			numArr = numArr[:len(numArr)-1]
		}
		if node.Right != nil {
			numArr = append(numArr, node.Val)
			dfs(node.Right)
			numArr = numArr[:len(numArr)-1]
		}
	}
	dfs(root)
	for _, ints := range resultArr {
		var tmpS string
		for i := 0; i < len(ints); i++ {
			tmpS += strconv.Itoa(ints[i])
			if i != len(ints)-1 {
				tmpS += "->"
			}
		}
		result = append(result, tmpS)

	}
	fmt.Println(resultArr)

	return result
}
