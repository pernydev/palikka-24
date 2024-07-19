package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/pernydev/palikka-24/app"
	"github.com/pernydev/palikka-24/app/connections"
	"github.com/pernydev/palikka-24/app/state"
)

func main() {
	godotenv.Load()
	fmt.Println("Starting Palikka backend server...")
	state.Load()
	go connections.StartDeltaTimer()
	go state.Autosave()
	app.InitServer()
}
