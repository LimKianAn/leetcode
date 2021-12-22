package main

func isIsomorphic(a, b string) bool {
	forward := map[byte]byte{} // from a to b
	back := map[byte]byte{} // from b to a

	for i := range a {
		x, y := a[i], b[i]
		if forward[x] != 0 && forward[x] != y || back[y] != 0 && back[y] != x {
			return false
		}
		forward[x] = y
		back[y] = x
	}
	return true
}
