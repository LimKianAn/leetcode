package main

func reverse0007(x int) int {
	tmp := 0
	for x != 0 {
		tmp = 10*tmp + x%10
		x /= 10
	}
	if tmp < -1<<31 || tmp > 1<<31-1 {
		return 0
	}
	return tmp
}
