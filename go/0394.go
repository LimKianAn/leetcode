package main

import (
	"log"
	"strconv"
)

func decodeString(s string) string {
	n := len(s)

	i := firstNumIndex(s)
	if i == n {
		return s // s consists of abc only
	}

	p := i + 1 // first possible [
	for s[p] != '[' {
		p++
	}

	q := p + 1 // finding the corresponding ]
	for count := 1; count > 0; {
		q++
		if s[q] == '[' {
			count++
		} else if s[q] == ']' {
			count--
		}
	}

	times, _ := strconv.Atoi(s[i:p])
	return s[:i] + multiply(decodeString(s[p+1:q]), times) + decodeString(s[q+1:])

}

func firstNumIndex(s string) (i int) { // s can start with abc
	for i < len(s) && s[i] > '9' {
		i++
	}
	return
}

func multiply(s string, times int) (multiplied string) {
	for i := 0; i < times; i++ {
		multiplied += s
	}
	return
}

func main() {
	log.Print(decodeString("3[a]2[bc]"))
}
