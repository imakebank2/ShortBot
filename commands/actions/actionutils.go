// Commonly used functions for commands can be added here

package actions

import (
	"github.com/bwmarrin/discordgo"
)

// Responds with text after an interaction
func interactionTextResponse(text string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: text,
		},
	})
}
