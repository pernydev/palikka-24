package state

import "fmt"

var State = make(map[string]uint8)

func PlaceBlock(x, y int, texture uint8) {
	if texture == 0 {
		delete(State, fmt.Sprint(x)+","+fmt.Sprint(y))
		return
	}
	State[fmt.Sprint(x)+","+fmt.Sprint(y)] = texture
}
