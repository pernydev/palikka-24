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
		AllowOrigins:     []string{"http://localhost:5173", "https://palikka-24.vercel.app"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/", routes.Index)
	router.GET("/api/socket", routes.Socket)
	router.POST("/api/place", routes.Place)
	router.GET("/api/grid", routes.Grid)
	router.POST("/api/execute", routes.Execute)

	router.Run(":" + os.Getenv("PORT"))
}
