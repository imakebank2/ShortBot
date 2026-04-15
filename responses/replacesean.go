package responses

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var letter_combinations = []string{
	"TION",
	"SION",
	"CION",
	"CIAN",
	"CEAN",
}

// Replaces specific letter combinations with "SEAN" in words and sends the modified message
func ReplaceSean(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.Fields(strings.ToUpper(m.Content))

	for _, word := range content {
		for _, combo := range letter_combinations {
			if strings.Contains(word, combo) {
				s.ChannelMessageSendReply(m.ChannelID, strings.Replace(word, combo, "**SEAN**", 1), m.Reference())
				break
			}
		}
	}
}
