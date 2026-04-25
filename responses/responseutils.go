package responses

import "github.com/bwmarrin/discordgo"

func replyToUserWithoutPing(text string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content:   text,
		Reference: m.Reference(), // Reply to the message
		AllowedMentions: &discordgo.MessageAllowedMentions{
			RepliedUser: false, // DO NOT PING
		},
	})
}
