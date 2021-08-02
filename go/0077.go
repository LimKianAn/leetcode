package main

func combine(n int, k int) (result [][]int) {
	do(&result, n, k, 1, []int{})
	return
}

func do(result *[][]int, n, k, head int, memo []int) {
	if len(memo) == k {
		alias := make([]int, k)
		copy(alias, memo)
		*result = append(*result, alias)
		return
	}

	for i := head; i <= n-(k-len(memo)-1); i++ {
		do(result, n, k, i+1, append(memo, i))
	}
	return
}
