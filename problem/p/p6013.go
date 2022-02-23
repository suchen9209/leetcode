package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head = new(ListNode)
	head.Val = 0

	var n1, n2, n3, n4, n5 = new(ListNode), new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	n1.Val = 3
	n2.Val = 1
	n3.Val = 0
	n4.Val = 4
	n5.Val = 0
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	fmt.Println(mergeNodes(head))
}

func mergeNodes(head *ListNode) *ListNode {
	tmpSum := 0
	returnHead := new(ListNode)
	tHead := returnHead

	for head.Next != nil {
		if head.Val == 0 {
			if tmpSum != 0 {
				returnHead.Val = tmpSum
				returnHead.Next = new(ListNode)
				returnHead = returnHead.Next
				tmpSum = 0
			}
		} else {
			tmpSum += head.Val
		}
		head = head.Next
	}
	if tmpSum > 0 {
		returnHead.Val = tmpSum
	}

	return tHead
}
