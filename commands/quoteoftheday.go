package commands

import (
	"log"
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

func QuoteOfTheDay(s *discordgo.Session, i *discordgo.InteractionCreate) {
	msgs, err := s.ChannelMessages("1490700344852873328", 100, "", "", "")

	if err != nil {
		log.Println(err)
		return
	}

	randomIndex := rand.IntN(len(msgs))
	quote := msgs[randomIndex]
	s.ChannelMessageSend(i.ChannelID, quote.Content)
}
