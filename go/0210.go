package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	numPr := make([]int, numCourses)     // numPr := number of prerequisites
	ccAfter := make([][]int, numCourses) // ccAfter := courses after a prerequisite

	for _, pair := range prerequisites {
		c, pr := pair[0], pair[1] // c := course, pr := prerequisite
		numPr[c]++
		ccAfter[pr] = append(ccAfter[pr], c)
	}

	prClearedCC := prerequisiteClearedCourses(numPr) // courses which can be taken freely
	for i := 0; i < len(prClearedCC); i++ { // prClearedCC grows; for-range doesn't grow
		cc := ccAfter[prClearedCC[i]]
		for _, c := range cc {
			if numPr[c]--; numPr[c] == 0 { // one of the prerequisite has been taken, but there might be still more prerequisites
				prClearedCC = append(prClearedCC, c)
			}
		}
	}

	if len(prClearedCC) == numCourses {
		return prClearedCC
	}
	return []int{}
}

func prerequisiteClearedCourses(numPr []int) (a []int) {
	for i := range numPr {
		if numPr[i] == 0 {
			a = append(a, i)
		}
	}
	return
}
