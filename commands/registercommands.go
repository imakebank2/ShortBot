package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/commands/actions"
)

var BotCommands = []Command{
	{"quoteoftheday", "Says a random quote", actions.QuoteOfTheDay},
	{"randomuser", "Gets a random user", actions.RandomUser},
}

func RegisterCommands(s *discordgo.Session) {
	for _, c := range BotCommands {
		c.Register(s)
	}
}
