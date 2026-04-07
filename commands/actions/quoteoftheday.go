package actions

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Says a random quote from shortgang-quotes channel
func QuoteOfTheDay(s *discordgo.Session, i *discordgo.InteractionCreate) {
	msgs, err := s.ChannelMessages("1490700344852873328", 100, "", "", "")

	if err != nil {
		log.Println(err)
		return
	}

	quote, err := randomElementFrom(msgs)

	if err != nil {
		log.Println(err)
		return
	}

	interactionTextResponse(quote.Content, s, i)
}
