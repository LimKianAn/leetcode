package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	m := middle(head)
	otherHead := m.Next
	m.Next = nil

	return merge(sortList(head), sortList(otherHead))
}

func merge(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val > b.Val {
		b.Next = merge(a, b.Next)
		return b
	}
	a.Next = merge(a.Next, b)
	return a
}

func middle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next      // step := 1 node
		fast = fast.Next.Next // step := 2 node
	}
	return slow
}
