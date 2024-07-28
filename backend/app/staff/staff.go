package staff

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pernydev/palikka-24/app/discolog"
	"github.com/pernydev/palikka-24/models"
	"github.com/pernydev/palikka-24/models/response"
)

func StaffOnly(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		discolog.Send(map[string]string{
			"error": "No token provided",
			"ip":    c.ClientIP(),
			"ua":    c.GetHeader("User-Agent"),
		}, "API ABUSE")

		response.Error("No token provided", 401, c)
		return false
	}

	data := models.AuthTokenData{}
	_, err := jwt.ParseWithClaims(token, &data, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		response.Error("Invalid token", 401, c)
		return false
	}

	if !data.Staff {
		discolog.Send(map[string]string{
			"error": "Unauthorized",
			"user":  data.ID,
			"ip":    c.ClientIP(),
			"ua":    c.GetHeader("User-Agent"),
		}, "API ABUSE")

		response.Error("Unauthorized", 403, c)
		return false
	}

	return true
}
