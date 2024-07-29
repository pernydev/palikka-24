package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pernydev/palikka-24/app/connections"
	"github.com/pernydev/palikka-24/app/discolog"
	"github.com/pernydev/palikka-24/app/state"
	"github.com/pernydev/palikka-24/app/state/cooldown"
	"github.com/pernydev/palikka-24/models"
	"github.com/pernydev/palikka-24/models/response"
)

func Place(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		response.Error(err.Error(), 500, c)
		return
	}

	fmt.Println("Place", int(data[0]), int(data[1]), int(data[2]))

	token := c.Request.Header.Get("Authorization")
	tokenData := models.AuthTokenData{}
	_, err = jwt.ParseWithClaims(token, &tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		response.Error("Invalid token", 401, c)
		return
	}

	if len(data) != 3 {
		discolog.Send(map[string]string{
			"error": "Invalid data packet. This endpoint only takes one place packet.",
			"user":  fmt.Sprintf("<@%s>", tokenData.ID),
			"ip":    c.ClientIP(),
			"ua":    c.GetHeader("User-Agent"),
		}, "API ABUSE")
		response.Error("Invalid data packet. This endpoint only takes one place packet.", 400, c)
		return
	}

	fmt.Println("Place", int(data[0]), int(data[1]), int(data[2]))

	if !state.Open {
		response.Error("The game is not open. Bro, we logged you, we know who you are, why you tryna API abuse?", 400, c)
		discolog.Send(map[string]string{
			"error": "Pixel placement attempted while the game is closed.",
			"user":  fmt.Sprintf("<@%s>", tokenData.ID),
			"ip":    c.ClientIP(),
			"ua":    c.GetHeader("User-Agent"),
		}, "API ABUSE")
		return
	}

	if cooldown.IsOnCooldown(c.ClientIP(), tokenData.ID) {
		response.Error("You are on cooldown. Please wait a bit before placing another pixel.", 420, c) // 420 blaze it
		discolog.Send(map[string]string{
			"error": "Pixel placement attempted while on cooldown.",
			"user":  fmt.Sprintf("<@%s>", tokenData.ID),
			"ip":    c.ClientIP(),
			"ua":    c.GetHeader("User-Agent"),
		}, "API ABUSE")
		return
	}

	connections.AddToDelta(data)

	state.PlaceBlock(int(data[0]), int(data[1]), uint8(data[2]))

	cooldown.SetCooldown(c.ClientIP(), tokenData.ID)

	response.Success(nil, c)
}
