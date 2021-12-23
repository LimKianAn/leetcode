package main

import "strings"

func isPalindrome0125(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
outer:
	for i < j {
		for !isValid(s[i]) {
			i++
			continue outer
		}
		for !isValid(s[j]) {
			j--
			continue outer
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isValid(b byte) bool {
	if ('a' <= b && b <= 'z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}
