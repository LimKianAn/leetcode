package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	extended := &ListNode{Next: head}
	pivot := extended
	for i := 0; i < left-1; i++ {
		pivot = pivot.Next
	}

	now := pivot.Next
	for i := 0; i < right-left; i++ {
		tmp := now.Next
		now.Next = now.Next.Next
		tmp.Next = pivot.Next
		pivot.Next = tmp
	}
	return extended.Next
}
