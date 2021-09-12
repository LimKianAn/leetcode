package main

func asteroidCollision(asteroids []int) (ii []int) {
	for _, v := range asteroids {
		if v < 0 {
			// ii becomes shorter
			for lastOneGoesRightward(ii) && -v > ii[len(ii)-1] {
				pop(&ii)
			}

			if lastOneGoesRightward(ii) {
				if -v == ii[len(ii)-1] {
					pop(&ii)
				}
			} else {
				push(&ii, v)
			}
		} else {
			push(&ii, v)
		}
	}
	return
}

func lastOneGoesRightward(ii []int) bool {
	return len(ii) > 0 && ii[len(ii)-1] > 0
}

func pop(ii *[]int) {
	(*ii) = (*ii)[:len(*ii)-1]
}

func push(ii *[]int, i int) {
	(*ii) = append((*ii), i)
}
