package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")

	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

}

// "Rewriting" it from Python to Go since it is way better for free/cheap hosting since it is much more performant + uses less storage and ram
// However, barrier to contribution is a bit higher so rip.
