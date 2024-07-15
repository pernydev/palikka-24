package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pernydev/palikka-24/app/routes"
)

func InitServer() {
	fmt.Println("Starting Palikka backend server...")

	router := gin.Default()
	router.GET("/", routes.Index)
	router.GET("/socket", routes.Socket)

	router.Run(":" + os.Getenv("PORT"))
}
