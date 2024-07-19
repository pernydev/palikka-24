package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pernydev/palikka-24/app/connections"
	"github.com/pernydev/palikka-24/app/state"
	"github.com/pernydev/palikka-24/models/response"
)

func Place(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		response.Error(err.Error(), 500, c)
		return
	}

	fmt.Println("Place", int(data[0]), int(data[1]), int(data[2]))

	if len(data) != 3 {
		response.Error("Invalid data packet. This endpoint only takes one place packet.", 400, c)
		return
	}

	connections.AddToDelta(data)
	state.PlaceBlock(int(data[0]), int(data[1]), uint8(data[2]))
	response.Success(nil, c)
}
