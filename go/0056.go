package main

// O(n log n) to O(n^2)
func merge0056(intervals [][]int) [][]int {
	Intervals(intervals).quickSort(0, len(intervals)-1)
	a := [][]int{intervals[0]} // a := answer
	i := 0
	for _, interval := range intervals {
		if curInterval := a[i]; curInterval[1] < interval[0] { // not overlapping
			a = append(a, interval)
			i++
		} else {
			curInterval[1] = max(curInterval[1], interval[1]) // extending
		}
	}
	return a
}

type Intervals [][]int

func (a Intervals) less(index, pivot int) bool {
	return a[index][0] < a[pivot][0] // comparing the start of an interval
}

// max is chosen to be pivot
func (a Intervals) partition(l, r int) int {
	curSmaller := l - 1
	for i := l; i < r; i++ {
		if a.less(i, r) {
			curSmaller++
			a.swap(curSmaller, i)
		}
	}
	curSmaller++
	a.swap(curSmaller, r)
	return curSmaller
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

// func main() {
// 	log.Print(merge([][]int{{4, 5}, {1, 4}, {0, 1}}))
// }
