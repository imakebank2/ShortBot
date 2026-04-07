package eventhandlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/commands"
)

// Called when someone messages in the server
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	commands.SendSup(s, m)
}
