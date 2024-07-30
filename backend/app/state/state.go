package state

import "fmt"

var State = make(map[string]uint8)
var Open = true
var End = false
var EndFinal = false
var Scene = 0
var Whited = false

func PlaceBlock(x, y int, texture uint8) {
	State[fmt.Sprint(x)+","+fmt.Sprint(y)] = texture
}

func GetBlock(x, y int) uint8 {
	return State[fmt.Sprint(x)+","+fmt.Sprint(y)]
}
