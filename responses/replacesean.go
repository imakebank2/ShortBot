package responses

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var letter_combinations = []string{
	"TION",
	"SSION",
	"CION",
	"CIAN",
	"CEAN",
}

// Replaces specific letter combinations with "SEAN" in words and sends the modified message
func ReplaceSean(s *discordgo.Session, m *discordgo.MessageCreate) {
	var repliedWords []string

	content := strings.Fields(strings.ToUpper(m.Content))

	for _, word := range content {
		for _, combo := range letter_combinations {
			if strings.Contains(word, combo) {
				repliedWords = append(repliedWords, strings.Replace(word, combo, "**SEAN**", 1))
			}
		}
	}

	if len(repliedWords) != 0 {
		replyToUserWithoutPing(strings.Join(repliedWords, ", "), s, m)
	}

}
