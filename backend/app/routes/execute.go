package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pernydev/palikka-24/app/connections"
	"github.com/pernydev/palikka-24/app/staff"
	"github.com/pernydev/palikka-24/app/state"
	"github.com/pernydev/palikka-24/models/response"
)

func Execute(c *gin.Context) {
	if !staff.StaffOnly(c) {
		return
	}

	data, err := c.GetRawData()
	if err != nil {
		response.Error(err.Error(), 500, c)
		return
	}

	connections.AddToDelta(data)
	for i := 0; i < len(data); i += 3 {
		state.PlaceBlock(int(data[i]), int(data[i+1]), uint8(data[i+2]))
	}

	for i := 0; i < len(data); i += 3 {
		x := data[i]
		y := data[i+1]
		texture := data[i+2]
		state.State[fmt.Sprint(x)+","+fmt.Sprint(y)] = texture
	}

	response.Success(nil, c)
}
