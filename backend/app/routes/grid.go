package routes

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pernydev/palikka-24/app/state"
)

func Grid(c *gin.Context) {
	data := make([]byte, 0)

	for coords, block := range state.State {
		coordsParsed := strings.Split(coords, ",")
		// parse as uint8
		x, _ := strconv.ParseUint(coordsParsed[0], 10, 8)
		y, _ := strconv.ParseUint(coordsParsed[1], 10, 8)

		data = append(data, byte(x))
		data = append(data, byte(y))
		data = append(data, block)

		fmt.Println("Grid", coordsParsed[0], coordsParsed[1], block)
		fmt.Println("Data", data)
	}

	c.Data(200, "application/octet-stream", data)
}
