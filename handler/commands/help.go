package commands

import (
	"fmt"
	"on-esports/utils"

	"github.com/bwmarrin/discordgo"
)

// Message is a generated ready to send message with description
// of each endpoint available.
type HelpCommand struct{
	Message string
}


// Constructor for /help command handler
// Takes a list of commands from main handler and
// generates a /help command message using formatter function call
// then returns an up-to-date /help command message with all available commands listed
func NewHelp(commands []*utils.Command) *HelpCommand {

	// formatting raw commands data to ready to send message
	title := "Privet, my name is Anatoliy 👋🏻\n\n.I love gaming. But more than anything, I love to check match statistics and I am here to help you to check out esport games.🎮\n\n\nAvailable commands:"
	message := utils.FormatHelp(commands)

	return &HelpCommand{
		Message: fmt.Sprintf("%s\n%s", title, message),
	}
}


// Sending generated /help command message to the chat
func (help *HelpCommand) HelpHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	s.ChannelMessageSend(m.ChannelID, help.Message)
}
