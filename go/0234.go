package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	for i, j := 0, len(vals)-1; i < j; {
		if vals[i] != vals[j] {
			return false
		}
		i++
		j--
	}
	return true
}
