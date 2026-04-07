package actions

import (
	"fmt"
	"math/rand/v2"

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

func randomElementFrom[T any](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, fmt.Errorf("slice is empty")
	}

	return s[rand.IntN(len(s))], nil
}
