package main

func numTrees(n int) int {
	memo := make([]int, n+1)
	memo[0], memo[1] = 1, 1
	for i := 2; i <= n; i++ { // i := number of nodes
		for j := 1; j <= i; j++ { // j := value of root
			// memo[j-1] = number of possible trees with the value range [1, j-1].
			// memo[i-j] = number of possible trees with the value range [j+1, i],
			// (continued) which equals one with the value range [1, i-j].
			// += because different i generate different numbers and we have to sum them up
			memo[i] += memo[j-1] * memo[i-j]
		}
	}
	return memo[n]
}
