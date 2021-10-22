package main

func findTheDifference(s string, t string) byte {
	n := len(s)
	ch := t[n]
	for i := 0; i < n; i++ {
		ch ^= s[i] ^ t[i]
	}
	return ch
}
