package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Description string
	Action      func(*discordgo.Session, *discordgo.InteractionCreate)
}

func (c Command) Register(s *discordgo.Session) {
	_, err := s.ApplicationCommandCreate(s.State.Application.ID, "",
		&discordgo.ApplicationCommand{
			Name:        c.Name,
			Description: c.Description,
		})

	if err != nil {
		log.Println(err)
	}
}
