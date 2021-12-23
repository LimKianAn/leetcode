package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1    // number of nodes
	p := head // p:= pointer
	for p.Next != nil {
		n++
		p = p.Next
	}

	k %= n
	if k == 0 {
		return head
	}

	// Connect the last node to the head
	p.Next = head

	// p is currently at the last node. Use the simplest case
	// to find the new end, ex 1, 2.
	for i := 0; i < (n - k); i++ {
		p = p.Next
	}

	// new head
	head = p.Next

	// new end
	p.Next = nil

	return head
}
