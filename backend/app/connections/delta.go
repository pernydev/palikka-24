package connections

import (
	"fmt"
	"time"
)

var delta [][]byte = make([][]byte, 0)

func AddToDelta(edit []byte) {
	delta = append(delta, edit)
}

func SendDelta() {
	for _, delta := range delta {
		fmt.Println("Sending delta")
		go SendToAll(delta)
	}

	delta = make([][]byte, 0)
}

func StartDeltaTimer() {
	for {
		SendDelta()
		<-time.After(1 * time.Second)
	}
}
