package commands

import (
	"on-esports/utils"

	"github.com/bwmarrin/discordgo"
)


// A list of quotes related to "ON Esports"
var onEsportsQuotes = []string{
	"Возьмите на работу плиииз, готов работать за спасибо :)",
	"ON Esports? More like ON Point Esports, am I right, comrades?",
	"Their blockchain tech is like a legendary loot box, you never know what epic esports goodies you'll find inside!",
	"They're revolutionizing esports faster than Anatoliy can down a bowl of borscht.",
	"Imagine combining the power of blockchain with the passion of esports, that's ON Esports in a nutshell (or should I say, a cybersphere?)",
	"They're like the ultimate esports wingman, helping players and fans level up their game.",
	"Using ON Esports feels like unlocking a secret level in your favorite game, full of new possibilities.",
	"They're the Willy Wonka of esports, creating a fantastical world where blockchain and competition collide.",
	"If you're an esports player or fan, ON Esports is like finding a rare and powerful artifact, it'll give you the edge you need.",
	"They're building the future of esports, brick by digital brick (or should I say, blockchain block?)",
	"Think of ON Esports as the ultimate esports hype machine, injecting pure excitement into every match and event.",
	"They're shaking up the esports scene like a well-timed EMP grenade, and it's glorious!",
	"If you're looking to be part of something revolutionary in esports, look no further than ON Esports.",
	"They're the champions of innovation, always pushing the boundaries of what's possible in esports.",
	"Imagine winning the esports lottery, that's the feeling you get when you discover ON Esports.",
	"Anatoliy guarantees ON Esports is more fun than a pixelated puppy, and that's saying something!",
}

// Creating a new /on-esports command
func NewOnEsports() *utils.Command {
	return &utils.Command{
		Command: "/on-esports",
		Description: "Gives you a quote about \"ON Esports\" company :)",
		Handler: OnEsportsHandler,
	}
}


// tha function that actually serves for /on-esports
func OnEsportsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	randomQuote := utils.RandomQuote(&onEsportsQuotes)

	formattedQuote := utils.FormatOnEsports(randomQuote)

	s.ChannelMessageSend(m.ChannelID, formattedQuote)
}
