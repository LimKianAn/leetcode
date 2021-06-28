package main

var vowels = [...]byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func reverseVowels(s string) string {
	bb := []byte(s)
	for i, j := 0, len(bb)-1; i < j; {
		if !isVowel(bb[i]) {
			i++
			continue
		}
		if !isVowel(bb[j]) {
			j--
			continue
		}
		bb[i], bb[j] = bb[j], bb[i]
		i++
		j--
	}
	return string(bb)
}

func isVowel(b byte) bool {
	for i := range vowels {
		if vowels[i] == b {
			return true
		}
	}
	return false
}
