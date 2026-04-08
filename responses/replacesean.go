package responses

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Replaces "tion" with "SEAN" in words and sends the modified message
func ReplaceSean(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Fields(strings.ToUpper(m.Content))

	for _, word := range content {
		if strings.Contains(word, "TION") {
			s.ChannelMessageSendReply(m.ChannelID, strings.ReplaceAll(word, "TION", "**SEAN**"), m.Reference())
		}
	}
}
