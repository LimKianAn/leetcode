package main

import "fmt"

func racecar_b(target int) int {
	type state struct {
		pos, speed int
	}
	q := []state{{0, 1}}
	isVisited := map[string]bool{
		"0,1":  true,
		"0,-1": true,
	}
	num := 0
	for len(q) > 0 {
		for qLen := len(q); qLen > 0; qLen-- {
			stt := q[0]
			q = q[1:]

			pos := stt.pos
			if pos == target {
				return num

			}

			speed := stt.speed
			{ // A
				newPos := pos + speed
				newSpeed := speed * 2
				key := fmt.Sprintf("%d,%d", newPos, newSpeed)
				if newPos > 0 && newPos < 2*target && !isVisited[key] {
					isVisited[key] = true
					q = append(q, state{newPos, newSpeed})
				}
			}

			{ // R
				newSpeed := 0
				if speed > 0 {
					newSpeed = -1
				} else {
					newSpeed = 1
				}
				key := fmt.Sprintf("%d,%d", pos, newSpeed)
				if pos > 0 && pos < 2*target && !isVisited[key] {
					isVisited[key] = true
					q = append(q, state{pos, newSpeed})
				}
			}
		}
		num++
	}
	return -1
}
