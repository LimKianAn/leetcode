package main

func minMutation(start string, end string, bank []string) int {
	geneStrSet := map[string]struct{}{}
	for _, geneStr := range bank {
		geneStrSet[geneStr] = struct{}{}
	}
	delete(geneStrSet, start) // never goes back to the start

	if _, ok := geneStrSet[end]; !ok {
		return -1
	}

	queue := []string{start}
	mutations := 0
	for len(queue) > 0 {
		qLen := len(queue)
		for qLen > 0 {
			geneStr := queue[0]
			queue = queue[1:]
			qLen--

			geneRunes := []rune(geneStr)
			for i := 0; i < len(geneRunes); i++ {
				originalRn := geneRunes[i] // records
				for _, rn := range []rune{'A', 'C', 'G', 'T'} {
					geneRunes[i] = rn
					if string(geneRunes) == end {
						return mutations + 1
					}

					if _, ok := geneStrSet[string(geneRunes)]; !ok {
						continue
					}
					queue = append(queue, string(geneRunes))
					delete(geneStrSet, string(geneRunes)) // not again
				}
				geneRunes[i] = originalRn // sets back
			}
		}
		mutations++ // next level
	}
	return -1
}
