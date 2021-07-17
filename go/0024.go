package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	e := &ListNode{Next: head}
	for p := e; p.Next != nil && p.Next.Next != nil; { // p := pointer
		n0 := p.Next
		n1 := p.Next.Next
		n2 := p.Next.Next.Next

		// Arrange the new order
		p.Next = n1
		n1.Next = n0
		n0.Next = n2

		// Set up the next round
		p = n0
	}
	return e.Next
}

// -1  0  1  2
//  x

// -1  1  0  2
//  x

// -1  1  0  2
//        x
