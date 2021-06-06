package main

var toOpening = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}
func isValid(s string) bool {
	openings := []rune{}
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			openings = append(openings, v)
		} else {
			if len(openings) == 0 {
				return false
			}
			last := len(openings) - 1
			if openings[last] != toOpening[v] {
				return false
			}
			openings = openings[:last]
		}
	}
	return len(openings) == 0
}
