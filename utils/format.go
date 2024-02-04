package utils

import (
	"fmt"
	"strings"
)

//makes the text look bold in Discord
func Bold(text string) string {

	return fmt.Sprintf("**%s**", text)
}


// makes text italic
func Italic(text string) string {

	return fmt.Sprintf("*%s*", text)
}

// makes text italic
func InlineCode(text string) string {

	return fmt.Sprintf("`%s`", text)
}


// Formats /quote command to ready-to-send format
// Gets just a string with a quote and adds the attributes as title and signature
func FormatQuote(quote string) string {
	
	title := "Quote of the day 💭:"
	signature := Italic("-- your lovely Anatoliy")
	quote = Bold(quote)

	return fmt.Sprintf("%s\n\n\"%s\"\n\n%s", title, quote, signature)
}

// Formats /on-esports command
// Adds signature for formatted quote on /on-esports command
func FormatOnEsports(quote string) string {

	signature := Italic("-- your lovely ON Esports 💙")
	quote = Bold(quote)

	return fmt.Sprintf("%s\n\n%s", quote, signature)
}


// Formats /help command
// gets raw commands information and merges them into one message
func FormatHelp(commands []*Command) string {
	
	var formattedCommands []string

	for _, command := range commands {
		commandName := command.Command
		description := Bold(command.Description)

		formattedCommand := fmt.Sprintf("`%s` -- %s", commandName, description)

		formattedCommands = append(formattedCommands, formattedCommand)
	}

	return strings.Join(formattedCommands, "\n\n")
}