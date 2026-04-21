package actions

import (
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

var OneVsOneOptions = []*discordgo.ApplicationCommandOption{
	{
		Type:        discordgo.ApplicationCommandOptionUser,
		Name:        "user",
		Description: "User to challenge",
		Required:    true,
	},
}

// 1v1 against another person (randomly decides who wins)
func OneVsOne(s *discordgo.Session, i *discordgo.InteractionCreate) {
	result := rand.IntN(2) // number from 0 - 1

	challengedMember := i.ApplicationCommandData().Options[0].UserValue(s)

	if result == 0 {
		interactionTextResponse(i.Member.DisplayName()+" won against "+challengedMember.DisplayName(), s, i)
	} else {
		interactionTextResponse(challengedMember.DisplayName()+" won against "+i.Member.DisplayName(), s, i)
	}
}
