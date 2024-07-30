package explosion

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pernydev/palikka-24/app/connections"
	"github.com/pernydev/palikka-24/app/state"
)

// The secret event at the end of the game, you can only place TNTs and they actually explode.

var indistructible = []int{
	1,
	216,
	218,
}

func TNT(x, y int) {
	placeTNT(x, y, false)
	<-time.After(700 * time.Millisecond)
	placeTNT(x, y, true)
	<-time.After(700 * time.Millisecond)
	placeTNT(x, y, false)
	<-time.After(700 * time.Millisecond)
	placeTNT(x, y, true)
	<-time.After(700 * time.Millisecond)
	placeTNT(x, y, false)
	<-time.After(700 * time.Millisecond)
	placeTNT(x, y, true)
	<-time.After(700 * time.Millisecond)
	breakArea(x, y)
}

func breakArea(x, y int) {
	data := []byte{}

	texture := 0
	if state.EndFinal {
		texture = 251
	}

	for i := -3; i <= 3; i++ {
		for j := -3; j <= 3; j++ {
			indistructibleBlock := false
			for _, block := range indistructible {
				if state.GetBlock(x+i, y+j) == uint8(block) {
					indistructibleBlock = true
					break
				}
			}
			if indistructibleBlock && !state.EndFinal {
				continue
			}

			random := rand.Float64()
			fmt.Println(random)

			if i*i+j*j <= 9 || random < 0.3 {
				data = append(data, byte(x+i))
				data = append(data, byte(y+j))
				data = append(data, byte(texture))
			}
		}
	}

	connections.AddToDelta(data)

	for i := 0; i < len(data); i += 3 {
		x := data[i]
		y := data[i+1]
		texture := data[i+2]
		state.PlaceBlock(int(x), int(y), uint8(texture))
	}
}

func placeTNT(x, y int, white bool) {
	fmt.Println("TNT placed at", x, y)
	data := []byte{byte(x), byte(y), 164}
	if white {
		data[2] = 0
	}

	connections.AddToDelta(data)
}
