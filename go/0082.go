package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	extended := &ListNode{Next: head}
	cur := extended // so that we are sure the next node has a different Val
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			v := cur.Next.Val
			cur.Next = cur.Next.Next.Next
			for cur.Next != nil && cur.Next.Val == v {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return extended.Next
}
