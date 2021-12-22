package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	next := addTwoNumbers(l1.Next, l2.Next)

	if sum >= 10 {
		sum %= 10
		next = addTwoNumbers(next, &ListNode{Val: 1})
	}

	return &ListNode{
		Val:  sum,
		Next: next,
	}
}
