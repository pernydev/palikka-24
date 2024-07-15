package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/pernydev/palikka-24/app"
)

func main() {
	godotenv.Load()
	fmt.Println("Starting Palikka backend server...")
	app.InitServer()
}
