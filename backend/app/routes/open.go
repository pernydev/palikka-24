package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pernydev/palikka-24/app/connections"
	"github.com/pernydev/palikka-24/app/staff"
	"github.com/pernydev/palikka-24/app/state"
)

func SetOpen(c *gin.Context) {
	if !staff.StaffOnly(c) {
		return
	}

	state.Open = !state.Open

	delta := byte(0)
	if state.Open {
		delta = 1
	}

	connections.SendToAll([]byte{delta})

	c.Status(204)
}

func GetOpen(c *gin.Context) {
	c.JSON(200, gin.H{
		"open": state.Open,
	})
}
