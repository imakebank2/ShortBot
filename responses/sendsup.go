package responses

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Says "sup" every time a message contains "up"
// Currently not used
func SendSup(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.ToLower(m.Content)

	// Check if it contains "up" anywhere
	if strings.Contains(content, "up") {
		s.ChannelMessageSendReply(m.ChannelID, "sup", m.Reference())
	}
}
