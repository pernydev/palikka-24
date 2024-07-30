package state

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pernydev/palikka-24/app/state/cooldown"
)

func Save(file *os.File) {
	newData := make([]byte, 0)
	for key, value := range State {
		coords := strings.Split(key, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		newData = append(newData, byte(x))
		newData = append(newData, byte(y))
		newData = append(newData, value)
	}
	file.WriteAt(newData, 0)
}

func Autosave() {
	time.Sleep(3 * time.Second)
	file, err := os.Create("state.palikka")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		time.Sleep(3 * time.Second)
		fmt.Println("Autosaving")
		Save(file)
		go cooldown.CleanUp() // Just run that every three seconds too
	}
}
