package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	eHead := head.Next
	o, e := head, eHead // o, e := odd, even
	for e != nil && e.Next != nil {
		o.Next = e.Next
		o = o.Next
		e.Next = o.Next
		e = e.Next
	}
	o.Next = eHead
	return head
}
