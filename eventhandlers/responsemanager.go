package eventhandlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/responses"
)

// Manages the bot's automatic (non-command) responses
func ReponseManager(s *discordgo.Session, m *discordgo.MessageCreate) {
	responses.SendSup(s, m)
}
