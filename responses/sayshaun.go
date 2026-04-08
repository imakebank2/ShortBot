package responses

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// SayShaun replaces "tion" with "SHAUN" in words and sends the modified message
func SayShaun(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.Fields(strings.ToUpper(m.Content))

	for _, word := range content {
		if strings.Contains(word, "TION") {
			s.ChannelMessageSendReply(m.ChannelID, strings.ReplaceAll(word, "TION", "**SHAUN**"), m.Reference())
		}
	}
}
