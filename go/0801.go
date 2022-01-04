package main

// func minSwap(a []int, b []int) int {
// 	ans := 1<<63 - 1
// 	var dfs func(i, count int)
// 	dfs = func(i, count int) {
// 		if count >= ans {
// 			return
// 		}

// 		if i == len(a) {
// 			ans = count
// 			return
// 		}

// 		if i == 0 || a[i-1] < a[i] && b[i-1] < b[i] { // always branches when i==0
// 			dfs(i+1, count)
// 		}

// 		if i == 0 || a[i-1] < b[i] && b[i-1] < a[i] { // always branches when i==0
// 			a[i], b[i] = b[i], a[i]
// 			dfs(i+1, count+1) // only this branch with a,b swapped
// 			a[i], b[i] = b[i], a[i]
// 		}
// 	}
// 	dfs(0, 0)
// 	return ans
// }

func minSwap(a []int, b []int) int {
	n := len(a)
	keep := make([]int, n)
	swap := make([]int, n)
	keep[0] = 0
	swap[0] = 1
	for i := 1; i < n; i++ {
		keep[i] = n + 1 // unreachable max val in case the following two cases don't hold
		swap[i] = n + 1
		if a[i-1] < a[i] && b[i-1] < b[i] {
			keep[i] = keep[i-1]
			swap[i] = swap[i-1] + 1
		}

		if a[i-1] < b[i] && b[i-1] < a[i] {
			keep[i] = min(keep[i], swap[i-1]) // keep[i] might have been set in the previous scenario
			swap[i] = min(swap[i], keep[i-1]+1)
		}
	}
	return min(keep[n-1], swap[n-1])
}
