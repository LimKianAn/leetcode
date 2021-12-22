package main

func countPrimes(n int) (count int) {
	a := make([]bool, n)
	for i := 2; i < n; i++ {
		if a[i] {
			continue
		}
		for j := 2 * i; j < n; j += i { // multiples of i
			a[j] = true
		}
		count++ // primes incremented here
	}
	return
}

// i 2,
// j 4,6,8,...

// i 3,
// j 6,9,12,...
