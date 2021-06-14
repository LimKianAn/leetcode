package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	i := len(a)
	j := len(b)
	carry := byte(0)
	sum := ""
	for i > 0 && j > 0 {
		tmp := a[i-1] - '0' + b[j-1] - '0' + carry // byte
		sum = fmt.Sprint(tmp%2) + sum
		carry = tmp >> 1
		i--
		j--
	}

	for i > 0 {
		tmp := a[i-1] - '0' + carry
		sum = fmt.Sprint(tmp%2) + sum
		carry = tmp >> 1
		i--
	}

	if carry == 1 {
		return "1" + sum
	}
	return sum
}
