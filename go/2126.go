package main

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	curr := uint64(mass)
	sort.Sort(sortableAsteroids(asteroids))
	for _, asasteroid := range asteroids {
		if curr < uint64(asasteroid) {
			return false
		}
		curr += uint64(asasteroid)
	}
	return true
}

type sortableAsteroids []int

func (s sortableAsteroids) Len() int {
	return len(s)
}

func (s sortableAsteroids) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableAsteroids) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	return
}
