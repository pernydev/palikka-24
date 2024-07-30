package app

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pernydev/palikka-24/app/routes"
)

func InitServer() {
	fmt.Println("Starting Palikka backend server...")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	router.GET("/", routes.Index)
	router.GET("/api/socket", routes.Socket)
	router.POST("/api/place", routes.Place)
	router.GET("/api/grid", routes.Grid)
	router.POST("/api/execute", routes.Execute)
	router.POST("/api/auth", routes.Auth)
	router.POST("/api/open", routes.SetOpen)
	router.GET("/api/open", routes.GetOpen)

	router.POST("/api/end", routes.SetEnd)
	router.GET("/api/end", routes.GetEnd)

	router.Run(":" + os.Getenv("PORT"))
}
