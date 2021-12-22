package main

func findTheDifference(s string, t string) byte {
	n := len(s)
	ch := t[n] // the last char in t, which is longer than n by 1
	for i := 0; i < n; i++ {
		ch ^= s[i] ^ t[i]
	}
	return ch
}
