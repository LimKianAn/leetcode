package main

import "strings"

func wordPattern(pattern string, s string) bool {
	wToB := map[string]byte{}   // word to byte
	bToW := map[byte]string{}   // byte to word
	ww := strings.Split(s, " ") // words
	if len(pattern) != len(ww) {
		return false
	}

	for i, w := range ww {
		b := pattern[i] // byte
		if wToB[w] > 0 && wToB[w] != b || bToW[b] != "" && bToW[b] != w {
			return false
		}
		wToB[w] = b
		bToW[b] = w
	}
	return true
}
