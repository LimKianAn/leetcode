package main

import "log"

// O(n log n) to O(n^2)
func merge(intervals [][]int) [][]int {
	Intervals(intervals).quickSort(0, len(intervals)-1)
	a := [][]int{intervals[0]} // a := answer
	i := 0
	for _, v := range intervals {
		if now := a[i]; v[0] > now[1] {
			a = append(a, v)
			i++
		} else {
			now[1] = max(now[1], v[1])
		}
	}
	return a
}

type Intervals [][]int

func (a Intervals) less(index, pivot int) bool {
	return a[index][0] < a[pivot][0]
}

func (a Intervals) partition(l, r int) int {
	i := l - 1
	for j := l; j < r; j++ {
		if a.less(j, r) {
			i++
			a.swap(i, j)
		}
	}
	a.swap(i+1, r)
	return i + 1
}

func (a Intervals) quickSort(l, r int) {
	if l >= r {
		return
	}
	p := a.partition(l, r) // p := pivot
	a.quickSort(l, p-1)
	a.quickSort(p+1, r)
}

func (a Intervals) swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	log.Print(merge([][]int{{4, 5}, {1, 4}, {0, 1}}))
}
