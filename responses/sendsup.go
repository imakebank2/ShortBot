package responses

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Says "sup" every time a message contains "sup"
func SendSup(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.ToLower(m.Content)

	// Check if it contains "sup" anywhere
	if strings.Contains(content, "sup") {
		s.ChannelMessageSend(m.ChannelID, "sup")
	}
}
