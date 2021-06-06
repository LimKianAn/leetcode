package main

func isPalindrome(x int) bool {
	tmp := x
	reverse := 0
	for tmp > 0 {
		reverse = 10*reverse + tmp%10
		tmp /= 10
	}
	return reverse == x
}
