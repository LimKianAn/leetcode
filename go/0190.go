package main

func reverseBits(num uint32) (a uint32) {
	for i := 0; i < 32; i++ {
		a <<= 1
		a += num & 1
		num >>= 1
	}
	return
}
