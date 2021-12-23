package main

func partition0086(head *ListNode, x int) *ListNode {
	h1 := &ListNode{} // h := head
	h2 := &ListNode{}
	n1 := h1 // n := node
	n2 := h2
	for head != nil {
		if head.Val < x {
			n1.Next = head
			n1 = n1.Next
		} else {
			n2.Next = head
			n2 = n2.Next
		}
		head = head.Next
	}
	n1.Next = h2.Next
	n2.Next = nil
	return h1.Next
}
