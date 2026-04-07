package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	// Events bot can receive (should be the minimum possible for least overhead)
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentMessageContent

	// session.AddHandler()

	if err = session.Open(); err != nil {
		log.Fatal(err)
	}

	defer log.Println("Bot offline.")
	defer session.Close()

	log.Println("Bot online! Press CTRL-C to exit.")

	// Test example
	//_, err = session.ChannelMessageSend("1490700697681924147", "aaaaa")

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// Blocks until CTRL-C or other signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// Test example: Sends "sup" whenever someone else says "sup"
func sendSup() {

}

// "Rewriting" it from Python to Go since it is way better for free/cheap hosting since it is much more performant + uses less storage and ram
// However, barrier to contribution is a bit higher so rip.
