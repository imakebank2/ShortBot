package actions

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/utils"
)

// Says a random user from the list (excluding apps)
func RandomUser(s *discordgo.Session, i *discordgo.InteractionCreate) {
	members, err := s.GuildMembers(i.GuildID, "", 1000)

	if err != nil {
		log.Println(err)
		return
	}

	var humanMembers []*discordgo.Member
	for _, m := range members {
		if !m.User.Bot {
			humanMembers = append(humanMembers, m)
		}
	}

	humanMember, err := utils.GetRandomElement(humanMembers)

	if err != nil {
		log.Println(err)
		return
	}

	interactionTextResponse(humanMember.User.DisplayName(), s, i)
}
