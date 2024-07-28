package snapshot

import (
	"fmt"
	"os"
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

	state.Save(f)
}

func Autosnapshot() {
	for {
		if state.Open {
			take()
		}
		time.Sleep(10 * time.Second)
	}
}
