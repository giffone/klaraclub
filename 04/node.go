package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}
}

func print(list *ListNode) {
	buf := list

	for buf != nil {
		fmt.Printf("%d -> ", buf.Val)
		buf = buf.Next
	}
}

func insert(list, node *ListNode) *ListNode {
	buf := list

	for buf != nil {
		if buf.Next == nil {
			buf.Next = node
			return list
		}
		buf = buf.Next
	}
	return list
}
