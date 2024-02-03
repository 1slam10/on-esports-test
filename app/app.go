package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"


	"github.com/bwmarrin/discordgo"
)

type App struct {
	Token string
}

func New(token string) *App {

	token = fmt.Sprintf("Bot %s", token)

	return &App{token}
}

func (app *App) Start() error {
	//creating a new Discrod session with our Discord Token
	sess, err := discordgo.New(app.Token)

	if err != nil {
		return err
	}

	sess.AddHandler(func (s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		message := m.Content

		s.ChannelMessageSend(m.ChannelID, message)
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()

	if err != nil {
		return err
	}

	fmt.Println("Anatoliy woke up and ready to serve!")

	defer sess.Close()


	// some code to make our session open and make the bot serve
	// channel waits for syscall such as interrupting from terminal: ^ + C as an example
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<- sc

	return nil
}