package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	e := &ListNode{Next: head} // extended head
	l, r := e, e               // left, right

	// r moves rightwards n times
	for i := 0; i < n; i++ {
		r = r.Next
	}

	for r.Next != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next

	return e.Next
}
