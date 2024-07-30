package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pernydev/palikka-24/app/connections"
	"github.com/pernydev/palikka-24/app/staff"
	"github.com/pernydev/palikka-24/app/state"
)

func GetEnd(c *gin.Context) {
	c.JSON(200, gin.H{
		"end": state.End,
	})
}

func SetEnd(c *gin.Context) {
	if !staff.StaffOnly(c) {
		return
	}

	if !state.Whited && state.EndFinal {
		newData := make([]byte, 0)
		for x := 0; x < 170; x++ {
			for y := 0; y < 67; y++ {
				data := []byte{byte(x), byte(y), 0}
				newData = append(newData, data...)
				state.PlaceBlock(x, y, 0)
			}
		}

		connections.AddToDelta(newData)

		state.Whited = true
		return
	}

	if state.EndFinal {
		if state.Scene == 8 {
			connections.SendToAll([]byte{2})
			return
		}
		cycleScene()
		return
	}

	if state.End {
		fmt.Println("EndFinal")
		state.EndFinal = true
	}

	state.End = true

	connections.SendToAll([]byte{3})
	c.Status(204)
}

func cycleScene() {
	state.Scene++

	// get the scene file from ./scenes/(state.Scene).palikka
	file := fmt.Sprintf("./scenes/%d.palikka", state.Scene)
	fmt.Println("Loading scene", file)
	fileContents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading scene file:", err)
		return
	}

	quick := false
	if state.Scene == 8 {
		quick = true
	}

	state.LoadScene(fileContents, quick)
}
