package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB // adds headB to a
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA // adds headA to b
		}
	}
	return a
}

/*
aaaaabbb
bbbaaaaa
     ^ hear onward both links must be the same if any
 */
