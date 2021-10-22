package main

func isPowerOfThree(n int) bool {
	for n >= 3 {
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}
	return n == 1
}
