package commands

import (
	"on-esports/utils"

	"github.com/bwmarrin/discordgo"
)


// AI generated list of playful quotes from Anatoliy
var funnyQuotes = []string{
	"Need stats faster than a pro gamer dodges bullets? Anatoliy got your back, comrade!",
	"Upcoming matches got you hyped like a free skin giveaway? Check the schedule, then tell Anatoliy your predictions... no lowball offers on victory predictions!",
	"Blockchain in esports? It's like putting rocket boosters on your grandma's knitting needles - unexpected, but it might just work!",
	"Practice like Anatoliy trains: sleep, eat borscht, dominate online! But for real stats and schedules, just ask.",
	"Anatoliy once beat a pro gamer in his sleep... probably. Just ask him about it, he might tell the (mostly true) story.",
	"Got more trash talk than a salty streamer after a loss? Save it for the arena, here we share stats and hype, not insults (unless they're funny insults, then Anatoliy might approve).",
	"Feeling lucky like you just found a golden weapon in a loot box? Check the upcoming matches, maybe your favorite team is about to pop off!",
	"Anatoliy loves esports more than a babushka loves pierogi, that's saying something! So trust him when he says these stats are fire.",
	"Want to be a pro gamer? Ask Anatoliy for tips... but be warned, his training involves questionable amounts of borscht and epic dance moves.",
	"Feeling lost in the esports world like a noob in a MOBA? Relax, Anatoliy is your friendly neighborhood guide, point him at the question and he'll blast you with info like a well-placed rocket launcher.",
	"Blockchain, esports, it's all a beautiful mystery to Anatoliy... but he knows where to find the best stats and schedules, that's not a mystery!",
	"Got a burning question about esports that would make even the Sphinx blush? Anatoliy might not have the answer, but he'll definitely make a hilarious joke about it.",
	"Remember, competition is good, but sportsmanship is better! Be like Anatoliy: win with grace, lose with laughter (and maybe a rematch request).",
	"Life's too short for boring esports stats. Anatoliy delivers them spicy, funny, and with a side of borscht-fueled enthusiasm.",
	"Don't underestimate the power of a passionate gamer with questionable internet access and an even more questionable sense of humor. That's Anatoliy, and he's here to make your esports experience epic.",
}


// Function to create a /quote command
func NewQuote() *utils.Command {
	return &utils.Command{
		Command: "/quote",
		Description: "Gives you a random quote by Anatoliy",
		Handler: QuoteHandler,
	}
}


// Message handler for /quote command to send random quote
func QuoteHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	randomQuote := utils.RandomQuote(&funnyQuotes)

	formattedQuote := utils.FormatQuote(randomQuote)

	s.ChannelMessageSend(m.ChannelID, formattedQuote)
}