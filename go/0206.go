package main

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var last *ListNode
	for head != nil{
		next := head.Next
		head.Next = last

		// Increment
		last = head
		head = next
	}
	return last
}