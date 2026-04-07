package responses

import "github.com/bwmarrin/discordgo"

// Says sup every time someone says sup
func SendSup(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "sup" {
		s.ChannelMessageSend(m.ChannelID, "sup")
	}
}
