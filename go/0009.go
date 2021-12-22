package main

func isPalindrome(x int) bool {
	reverse := 0
	for tmp := x; tmp > 0; {
		reverse = 10*reverse + tmp%10
		tmp /= 10
	}
	return reverse == x
}
