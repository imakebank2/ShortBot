package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/imakebank2/shortbot/eventhandlers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	// Add event handlers here:
	session.AddHandler(eventhandlers.MessageCreate)

	// Events bot can receive (should be the minimum possible for least overhead)
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	if err = session.Open(); err != nil {
		log.Fatal(err)
	}

	defer log.Println("Bot offline.")
	defer session.Close()

	log.Println("Bot online! Press CTRL-C to exit.")

	botChannelID := "1490700697681924147" // Channel ID of bots-arena
	session.ChannelMessageSend(botChannelID, "I am online!")

	// Blocks until CTRL-C or other signal is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
