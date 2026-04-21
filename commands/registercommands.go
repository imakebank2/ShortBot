package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/commands/actions"
)

// Add your commands here
var BotCommands = []Command{
	{"quoteoftheday", "Says a random quote from #shortgang-quotes", actions.QuoteOfTheDay, nil},
	{"randomuser", "Gets a random user", actions.RandomUser, nil},
	{"67", "Info about the 67.", actions.SixEvenInfo, nil},
	{"1v1", "1v1 someone", actions.OneVsOne, actions.OneVsOneOptions},
}

func RegisterCommands(s *discordgo.Session) {
	for _, c := range BotCommands {
		c.Register(s)
	}
}
