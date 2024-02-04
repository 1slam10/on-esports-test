package utils

import (
	"github.com/bwmarrin/discordgo"
)


// Structure for commands to make /help command generate the help message with all commands available
// Command is the name of the command itself e.g. "/ping", "/privet"
// Description is a short description what this command actually do
// Handler is a function that gets pointers to channel and message creator, and then serves the 'Command' endpoint
type Command struct {
	Command string
	Description string
	Handler func (*discordgo.Session, *discordgo.MessageCreate)
}