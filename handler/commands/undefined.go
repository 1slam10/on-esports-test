package commands

import (
	"fmt"

	"on-esports/utils"

	"github.com/bwmarrin/discordgo"
)


// Function to handle an undefinded commands
// For the cases when user sent a wrong command through slash
// Formats output using utility formatter functions and send an alert that command sent by user is unknown
func UndefinedHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	title := utils.Bold("Nichego neponyatno!!!")
	help := utils.InlineCode("help")

	undefinedMessage := fmt.Sprintf("%s\n\n⛔️The command is not defined⛔️\n\nPlease send %s to stay up to date with Anatoliy's commands", title, help)

	s.ChannelMessageSend(m.ChannelID, undefinedMessage)
}
