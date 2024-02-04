package main

import (
	"github.com/joho/godotenv"
	"log"
	"on-esports/app"
	"os"
)

func main() {
	// Loading discord bot access token from env file into current process
	godotenv.Load()
	token := os.Getenv("DISCORD_TOKEN")

	app := app.New(token)

	err := app.Start()

	//checking if any error occured during app creation and running process
	if err != nil {
		log.Fatal(err)
	}
}
