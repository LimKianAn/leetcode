package main

import "strings"

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for i, s := range ss {
		ss[i] = reverse(s)
	}
	return strings.Join(ss, " ")
}

func reverse(s string) string {
	bb := []byte(s)
	for i, j := 0, len(bb)-1; i < j; {
		bb[i], bb[j] = bb[j], bb[i]
		i++
		j--
	}
	return string(bb)
}
