package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// We don't actually delete the node, but rather change its value to the next node's
// and skip the next node.
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
