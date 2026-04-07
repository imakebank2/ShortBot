package commands

import (
	"github.com/bwmarrin/discordgo"
)

var BotCommands = []Command{
	{"quoteoftheday", "Says a random quote", QuoteOfTheDay},
}

func RegisterCommands(s *discordgo.Session) {
	for _, c := range BotCommands {
		c.Make(s)
	}
}
