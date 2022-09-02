package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeSwap() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	l = swapPairs(l)
	print(l)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	second := head.Next
	head.Next = swapPairs(second.Next)
	second.Next = head
	return second
}

func print(head *ListNode) {
	buf := head

	for buf != nil {
		fmt.Printf("%d -> ", buf.Val)
		buf = buf.Next
	}
}
