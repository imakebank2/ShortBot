package eventhandlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/responses"
)

// Add your responses here
var botResponses = []func(*discordgo.Session, *discordgo.MessageCreate){
	responses.ReplaceSean,
	responses.SixSevenEvent,
}

// Manages the bot's automatic (non-command) responses
// Executes every time a message is sent
func ResponseManager(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, response := range botResponses {
		response(s, m)
	}
}
