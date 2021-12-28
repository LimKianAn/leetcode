package main

func findCircleNumDFS(isConnected [][]int) int {
	n := len(isConnected)
	var clearConnectedNodes func(i int)
	clearConnectedNodes = func(i int) {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 0 {
				continue
			}
			isConnected[i][j] = 0 // node j is connected; clear the connection
			isConnected[j][i] = 0 // the corresponding place
			clearConnectedNodes(j)
		}
	}

	num := 0
	for i := 0; i < n; i++ {
		if isConnected[i][i] == 0 { // connected with another node
			continue
		}
		num++
		clearConnectedNodes(i)
	}
	return num
}
