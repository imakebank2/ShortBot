package actions

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func TungTungTungSahur(s *discordgo.Session, i *discordgo.InteractionCreate) {
	interactionTextResponse(strings.Repeat("tung ", 100), s, i)
	s.ChannelMessageSend(i.ChannelID,
		"https://tenor.com/view/tung-tungtung-tungtungtung-sahur-tungtungtungsahur-tungtungsahur-gif-6699270143817937548",
	)
}
