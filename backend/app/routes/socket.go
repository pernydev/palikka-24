package routes

import (
	"fmt"
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

	go func() {
		for msg := range send {
			fmt.Println("Sending message to", uuid)
			err := conn.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				fmt.Println("error sending message:", err)
				connections.RemoveConnection(uuid)
				conn.Close()
				return
			}
		}
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error reading message:", err)
			break
		}
	}
}
