package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-2] + memo[i-1]
	}
	return memo[n]
}
