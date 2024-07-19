package state

import "fmt"

var State = make(map[string]uint8)

func PlaceBlock(x, y int, texture uint8) {
	State[fmt.Sprint(x)+","+fmt.Sprint(y)] = texture
}