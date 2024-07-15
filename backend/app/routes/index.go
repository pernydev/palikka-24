package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pernydev/palikka-24/models/response"
)

func Index(c *gin.Context) {
	response.Success("Hello, world!", c)
}
