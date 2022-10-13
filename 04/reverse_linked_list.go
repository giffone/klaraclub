package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	new := reverseList(head.Next)
	head.Next = nil // cut
	return insert(new, head)
}
