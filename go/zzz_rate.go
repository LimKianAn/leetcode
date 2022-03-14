package main

var rates map[string]map[string]float64

func setRate(inputCurr, outputCurr string, rate float64) {
	if rates[inputCurr] == nil {
		rates[inputCurr] = map[string]float64{}
	}
	if rates[outputCurr] == nil {
		rates[outputCurr] = map[string]float64{}
	}
	rates[inputCurr][outputCurr] = rate
	rates[outputCurr][inputCurr] = 1 / rate
}

func getRate(inputCurr, outputCurr string) float64 {
	var dfs func(x, z string, visited map[string]bool) float64
	dfs = func(x, z string, visited map[string]bool) float64 {
		if x == z {
			return 1
		}

		visited[x] = true
		for y, rateXY := range rates[x] {
			if visited[y] {
				continue
			}

			rateYZ := dfs(y, z, visited)
			if rateYZ == -1 {
				continue
			}
			return rateXY * rateYZ
		}
		return -1
	}
	return dfs(inputCurr, outputCurr, map[string]bool{})
}
