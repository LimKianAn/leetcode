// 2022.02.07

package main

func groupStrings(words []string) []int {
	m := map[int]int{}
	for _, word := range words {
		bitField := 0
		for _, rn := range word {
			bitField |= 1 << (rn - 'a')
		}
		m[bitField]++
	}

	var dfs func(bitfield int) int // returns bitfield's group size
	dfs = func(bitfield int) int {
		if m[bitfield] == 0 {
			return 0
		}
		ans := m[bitfield]
		delete(m, bitfield)

		for i := 0; i < 26; i++ {
			ans += dfs(bitfield ^ 1<<i) // by adding or deleting a letter
			for j := i + 1; j < 26; j++ {
				if bitfield>>i&1 != bitfield>>j&1 { // either i or j is set
					ans += dfs(bitfield ^ 1<<i ^ 1<<j) // by adding one letter and deleting the other
				}
			}
		}
		return ans
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	groupNum, maxGroupSize := 0, 0
	for len(m) > 0 {
		groupNum++
		for k := range m { // only the first key-value pair
			maxGroupSize = max(maxGroupSize, dfs(k))
			break
		}
	}
	return []int{groupNum, maxGroupSize}
}
