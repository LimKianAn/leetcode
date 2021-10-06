package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

// (x0+x1+x2+x1)/2=x0+x1
// x0+2*x1+x2=2*x0+2*x1
// x2=x0
