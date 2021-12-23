package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	coursesAfterPrerequisite := make([][]int, numCourses)
	numPrerequisites := make([]int, numCourses)

	for _, pair := range prerequisites {
		course, prerequisite := pair[0], pair[1]
		numPrerequisites[course]++
		coursesAfterPrerequisite[prerequisite] = append(coursesAfterPrerequisite[prerequisite], course)
	}

	prerequisiteClearedCourses := prerequisiteClearedCourses(numPrerequisites)
	for i := 0; i < len(prerequisiteClearedCourses); i++ { // for-range doesn't grow
		cc := coursesAfterPrerequisite[prerequisiteClearedCourses[i]]
		for _, c := range cc {
			if numPrerequisites[c]--; numPrerequisites[c] == 0 { // one of the prerequisites taken
				prerequisiteClearedCourses = append(prerequisiteClearedCourses, c)
			}
		}
	}

	return len(prerequisiteClearedCourses) == numCourses
}

func prerequisiteClearedCourses0207(numPrerequisites []int) (a []int) {
	for i := range numPrerequisites {
		if numPrerequisites[i] == 0 {
			a = append(a, i)
		}
	}
	return
}

// func main() {
// 	log.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
// }
