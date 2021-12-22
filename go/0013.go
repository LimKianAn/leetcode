package main

var integer = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	last := len(s) - 1
	sum := integer[s[last]]
	for i := last; i > 0; i-- {
		left := integer[s[i-1]]
		right := integer[s[i]]
		if left < right {
			sum -= left
		} else {
			sum += left
		}
	}
	return sum
}
