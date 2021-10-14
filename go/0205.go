package main

func isIsomorphic(a, b string) bool {
	forward := map[byte]byte{}
	back := map[byte]byte{}

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
