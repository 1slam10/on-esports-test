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

// The initial project should have been more complex and more creative than what I have now. 
// The features that I wanted to add initially listed below

// TODO: Add command to check latest news of given discipline
// TODO: Add command to check given player’s avg. statistics
// TODO: Add command to check given player's performance on given match
// TODO: Add command to check past game result of given team in the given discipline
// TODO: Add command to get upcoming games and tournaments of specified team or discipline.
// TODO: Add command to get H2H (Head-to-Head) stats of two commands
// TODO: Add "gamification": to implement coins
// TODO: Add command to create a bet on specified match using those coins
// TODO: Add command to get user's rating based on their balance / bet winning rate etc.

// Was very busy by stressing about that I don't understand anything, but in fact, it was much easier than expected.
// (I don't mean it is easy, code quality is not the best tho, could be better, but it is what it is, try to make it sexy).
// Anyways, there are a lot if things to do that I could do most earlier