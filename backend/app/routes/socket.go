package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pernydev/palikka-24/app/connections"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Socket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	send := make(chan []byte)
	uuid := uuid.New().String()

	connections.AddConnection(uuid, &connections.Connection{
		UUID:        uuid,
		SendChannel: send,
	})

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		send <- msg
	}
}
