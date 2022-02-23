// 2022.02.07

package main

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	n := len(s)

	p := 1
	for i := 1; i < k; i++ { // k-1 times
		p = p * power % modulo
	} // p=power^(k-1)%modulo

	val := func(i int) int {
		return int(s[i]-'a') + 1
	}

	ans := 0
	cur := 0
	for i := n - 1; i >= 0; i-- {
		if i+k < n {
			cur = cur + modulo - val(i+k)*p%modulo // + module so that no negative value would occur
		}
		cur = (cur*power + val(i)) % modulo

		if i+k <= n && cur == hashValue {
			ans = i
		}
	}
	return s[ans : ans+k]
}
