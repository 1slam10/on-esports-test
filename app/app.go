package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"on-esports/handlers"

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

	// adding command handlers to our session
	sess.AddHandler(handlers.OnEsportsHandler)
	sess.AddHandler(handlers.QuoteHandler)

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