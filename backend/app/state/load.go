package state

import (
	"fmt"
	"os"
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

		State[fmt.Sprint(x)+","+fmt.Sprint(y)] = texture
		fmt.Println("Loaded", x, y, texture)
	}
}
