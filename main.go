package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/imakebank2/shortbot/handlers"

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

	// Add handlers here:
	session.AddHandler(handlers.SendSup)

	// Events bot can receive (should be the minimum possible for least overhead)
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentMessageContent

	if err = session.Open(); err != nil {
		log.Fatal(err)
	}

	defer log.Println("Bot offline.")
	defer session.Close()

	log.Println("Bot online! Press CTRL-C to exit.")

	// botChannelID := "1490700697681924147"
	// session.ChannelMessageSend(botChannelID, "I am online!")

	// Blocks until CTRL-C or other signal is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

// "Rewriting" it from Python to Go since it is way better for free/cheap hosting since it is much more performant + uses less storage and ram
// However, barrier to contribution is a bit higher so rip.
