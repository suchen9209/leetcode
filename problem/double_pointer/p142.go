package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head = new(ListNode)
	head.Val = 3

	var n1, n2, n3 = new(ListNode), new(ListNode), new(ListNode)
	n1.Val = 2
	n2.Val = 0
	n3.Val = -4
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	fmt.Println(detectCycle(head))
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			tmp := head
			for tmp != slow {
				tmp = tmp.Next
				slow = slow.Next
			}
			return tmp
		}
	}

	return nil
}
