package main

// func longestMountain(arr []int) int {
// 	n := len(arr)
// 	up := make([]int, n)
// 	for i := 1; i < n; i++ {
// 		if arr[i-1] < arr[i] {
// 			up[i] = up[i-1] + 1
// 		}
// 	}

// 	down := make([]int, n)
// 	for i := n - 2; i >= 0; i-- {
// 		if arr[i] > arr[i+1] {
// 			down[i] = down[i+1] + 1
// 		}
// 	}

// 	ans := 0
// 	for i := 0; i < n; i++ {
// 		if up[i] != 0 && down[i] != 0 {
// 			ans = max(ans, up[i]+1+down[i])
// 		}
// 	}

// 	if ans < 3 {
// 		return 0
// 	}
// 	return ans
// }

func longestMountain(arr []int) int {
	up, down := 0, 0
	ans := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			if down > 0 {
				up, down = 0, 0
			}
			up++
		} else if arr[i-1] > arr[i] {
			down++
		} else {
			up, down = 0, 0
		}

		if up > 0 && down > 0 {
			ans = max(ans, up+1+down)
		}
	}

	if ans < 3 {
		return 0
	}
	return ans
}
