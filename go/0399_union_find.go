package main

func calcEquationUnionFind(equations [][]string, values []float64, queries [][]string) []float64 {
	type root struct {
		id    string
		ratio float64
	}
	rootMap := map[string]*root{}

	var getRoot func(id string) *root
	getRoot = func(id string) *root {
		if rootID := rootMap[id].id; rootID != id {
			newRoot := getRoot(rootID)
			rootMap[id].id = newRoot.id
			rootMap[id].ratio *= newRoot.ratio
		}
		return rootMap[id]
	}

	for i, equation := range equations {
		a := equation[0]
		b := equation[1]
		ratio := values[i]
		if rootMap[a] == nil && rootMap[b] == nil {
			rootMap[a] = &root{b, ratio} // just chooses one as the root
			rootMap[b] = &root{b, 1}
		} else if rootMap[a] == nil {
			rootMap[a] = &root{b, ratio}
		} else if rootMap[b] == nil {
			rootMap[b] = &root{a, 1 / ratio}
		} else {
			rootA := getRoot(a)
			rootB := getRoot(b)
			if rootA.id != rootB.id { // two distinct graphs
				rootMap[rootA.id] = &root{rootB.id, ratio / rootA.ratio * rootB.ratio} // merges rootA into rootB; the ratio path rootA -> a -> b -> rootB
			}
		}
	}

	ans := []float64{}
	for _, query := range queries {
		a := query[0]
		b := query[1]
		if rootMap[a] == nil || rootMap[b] == nil {
			ans = append(ans, -1)
			continue
		}

		rootA := getRoot(a)
		rootB := getRoot(b)
		if rootA.id != rootB.id {
			ans = append(ans, -1)
		} else {
			ans = append(ans, rootA.ratio/rootB.ratio) // the ratio path a -> rootA (rootB) -> b
		}
	}
	return ans
}
