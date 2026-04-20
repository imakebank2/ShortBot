package actions

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/utils"
)

// Says a random quote from shortgang-quotes channel
func QuoteOfTheDay(s *discordgo.Session, i *discordgo.InteractionCreate) {

	channels, err := s.GuildChannels(i.GuildID)

	if err != nil {
		log.Println(err)
		return
	}

	var channelID string

	for _, c := range channels {
		if c.Name == "shortgang-quotes" {
			channelID = c.ID
			break
		}
	}

	msgs, err := s.ChannelMessages(channelID, 100, "", "", "")

	if err != nil {
		log.Println(err)
		return
	}

	quote, err := utils.GetRandomElement(msgs)

	if err != nil {
		log.Println(err)
		return
	}

	interactionTextResponse(quote.Content, s, i)
}
