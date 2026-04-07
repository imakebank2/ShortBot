package eventhandlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/commands"
)

// Manages slash commands
func CommandManager(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	for _, c := range commands.BotCommands {
		if c.Name == i.ApplicationCommandData().Name {
			c.Action(s, i)
		}
	}
}
