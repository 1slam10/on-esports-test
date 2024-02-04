package handler

import (
	"strings"

	"on-esports/utils"
	"on-esports/handler/commands"

	"github.com/bwmarrin/discordgo"
)


// Main handler to handle any message input from user
// Commands is a map of commands. The key is a command e.g. '/quote', '/pivo' etc.
// CommandList stores commands as a slice to make /help command be able to get generated
// Help is a special command that stores a help info about all available commands added to the handler
type Handler struct {
	Commands map[string]*utils.Command
	CommandsList []*utils.Command
	Help *commands.HelpCommand
}



func NewHandler() Handler {
	//creating a new message handler
	handler := Handler{
		Commands: make(map[string]*utils.Command),
	}

	// creating instances for each of the commands
	onEsports := commands.NewOnEsports()
	quote := commands.NewQuote()

	
	// adding commands to the handler
	handler.AddCommand(onEsports)
	handler.AddCommand(quote)

	// creating a /help command separately from the rest
	// because help should process other commands info only after they've been added
	help := commands.NewHelp(handler.CommandsList)
	handler.Help = help

	return handler
}


// The main method of the Handler
// This is where the magic begins and dreams come true (just kidding)
// The method to handle all the incoming messages and commands
func (h *Handler) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
		
	// Ignoring messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := m.Content

	// checking if the message sent by user
	// is either a command or not
	if !strings.HasPrefix(message, "/") {
		return
	}

	// trimming all unnecessary whitespaces and excess symbols
	command := strings.Fields(message)[0]

	if command == "/help" {
		h.Help.HelpHandler(s, m)
		return
	}

	// Retrieving the respective handle for a specific key
	// using second bool argument to check if the command entered by user
	// is available for our bot
	handler, defined := h.Commands[command]

	// if there is no such a command it sends message about unknown command
	if !defined {
		commands.UndefinedHandler(s, m)
		return
	}

	// If all the checks passed, just handling the command
	// with its own handler
	handler.Handler(s, m)
}


// Function to "register" commands and their handlers in the main handler
func (h *Handler) AddCommand(command *utils.Command) {
	h.Commands[command.Command] = command
	h.CommandsList = append(h.CommandsList, command)
}




