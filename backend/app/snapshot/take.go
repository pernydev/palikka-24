package snapshot

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pernydev/palikka-24/app/state"
)

func take() {
	time := time.Now().Unix()
	filePath := fmt.Sprintf("snapshots/%d.palikka", time)

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	newData := make([]byte, 0)
	for key, value := range state.State {
		coords := strings.Split(key, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		newData = append(newData, byte(x))
		newData = append(newData, byte(y))
		newData = append(newData, value)
	}
	f.WriteAt(newData, 0)
}

func Autosnapshot() {
	for {
		if state.Open {
			take()
		}
		time.Sleep(10 * time.Second)
	}
}
