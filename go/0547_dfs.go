package main

func findCircleNum_d(isConnected [][]int) int {
	var clearConnectedNodes func(i int)
	clearConnectedNodes = func(i int) {
		for j := range isConnected {
			if isConnected[i][j] == 0 {
				continue
			}
			isConnected[i][j] = 0 // node j is connected; clear the connection
			isConnected[j][i] = 0 // the corresponding place
			clearConnectedNodes(j)
		}
	}

	ans := 0
	for i := range isConnected {
		if isConnected[i][i] == 0 { // connected with another node and cleared in the previous round
			continue
		}
		ans++
		clearConnectedNodes(i)
	}
	return ans
}
