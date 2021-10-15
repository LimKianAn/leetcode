package main

import "log"

func canFinish(numCourses int, prerequisites [][]int) bool {
	numPr := make([]int, numCourses)     // numPr := number of prerequisites
	ccAfter := make([][]int, numCourses) // ccAfter := courses after a prerequisite

	for _, pair := range prerequisites {
		c, pr := pair[0], pair[1] // c := course, pr := prerequisite
		numPr[c]++
		ccAfter[pr] = append(ccAfter[pr], c)
	}

	prClearedCC := prerequisiteClearedCourses(numPr)
	for i := 0; i < len(prClearedCC); i++ { // prClearedCC grows; for-range doesn't grow
		cc := ccAfter[prClearedCC[i]]
		for _, c := range cc {
			if numPr[c]--; numPr[c] == 0 {
				prClearedCC = append(prClearedCC, c)
			}
		}
	}

	return len(prClearedCC) == numCourses
}

func prerequisiteClearedCourses(numPr []int) (a []int) {
	for i := range numPr {
		if numPr[i] == 0 {
			a = append(a, i)
		}
	}
	return
}

func main() {
	log.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}
