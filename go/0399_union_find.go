// 2022.01.17

package main

func calcEquation_u(equations [][]string, values []float64, queries [][]string) []float64 {
	type ancestor struct {
		name  string
		ratio float64
	}
	ancestorMap := map[string]*ancestor{} // a/b = ratioAB; a -> (b, ratioAB)

	var findRoot func(name string) *ancestor
	findRoot = func(name string) *ancestor {
		ancestor := ancestorMap[name]
		if ancestor.name != name {
			rt := findRoot(ancestor.name)
			ancestor.name = rt.name // path compression
			ancestor.ratio *= rt.ratio
		}
		return ancestorMap[name]
	}

	for i, e := range equations {
		a := e[0]
		b := e[1]
		ratioAB := values[i]
		if ancestorMap[a] == nil && ancestorMap[b] == nil {
			ancestorMap[a] = &ancestor{b, ratioAB} // just chooses one as the root
			ancestorMap[b] = &ancestor{b, 1}
		} else if ancestorMap[a] == nil {
			ancestorMap[a] = &ancestor{b, ratioAB}
		} else if ancestorMap[b] == nil {
			ancestorMap[b] = &ancestor{a, 1 / ratioAB}
		} else {
			rtA := findRoot(a)
			rtB := findRoot(b)
			if rtA.name != rtB.name { // two distinct graphs
				ancestorMap[rtA.name] = &ancestor{rtB.name, 1 / rtA.ratio * ratioAB * rtB.ratio} // merges rootA into rootB; the ratio path rootA -> a -> b -> rootB
			}
		}
	}

	ans := []float64{}
	for _, qr := range queries {
		a := qr[0]
		b := qr[1]
		if ancestorMap[a] == nil || ancestorMap[b] == nil {
			ans = append(ans, -1)
			continue
		}

		rtA := findRoot(a) // updates ancestorMap!
		rtB := findRoot(b)
		if rtA.name != rtB.name {
			ans = append(ans, -1)
		} else {
			ans = append(ans, rtA.ratio/rtB.ratio) // the ratio path a -> rootA (rootB) -> b
		}
	}
	return ans
}
