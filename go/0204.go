package main

func countPrimes(n int) (count int) {
	a := make([]bool, n)
	for i := 2; i < n; i++ {
		if a[i] {
			continue
		}
		for j := 2 * i; j < n; j += i {
			a[j] = true
		}
		count++
	}
	return
}
