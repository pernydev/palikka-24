package state

import (
	"fmt"
	"os"
	"time"

	"github.com/pernydev/palikka-24/app/connections"
)

func Load() {
	f, err := os.Open("state.palikka")
	if err != nil {
		return
	}
	defer f.Close()

	data := make([]byte, 0)

	for {
		b := make([]byte, 3)
		_, err := f.Read(b)
		if err != nil {
			break
		}

		data = append(data, b...)
	}

	fmt.Println("Loaded", len(data), "bytes")

	for i := 0; i < len(data); i += 3 {
		x := data[i]
		y := data[i+1]
		texture := data[i+2]

		if texture == 0 {
			continue
		}

		State[fmt.Sprint(x)+","+fmt.Sprint(y)] = texture
		fmt.Println("Loaded", x, y, texture)
	}
}

func LoadScene(data []byte, quick bool) {

	for i := 0; i < len(data); i += 3 {
		x := data[i]
		y := data[i+1]
		texture := data[i+2]

		fmt.Println("Loaded", x, y, texture)

		State[fmt.Sprint(x)+","+fmt.Sprint(y)] = texture
		connections.AddToDelta([]byte{x, y, texture})
		if !quick {
			<-time.After(10 * time.Millisecond)
		}
	}
}
