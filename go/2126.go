// 2022.02.25

package main

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, asasteroid := range asteroids {
		if mass < asasteroid {
			return false
		}
		mass += asasteroid
	}
	return true
}
