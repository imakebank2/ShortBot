package actions

import "github.com/bwmarrin/discordgo"

var SetBirthdayOptions = []*discordgo.ApplicationCommandOption{
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "month",
		Description: "Birthday Month",
		Required:    true,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "day",
		Description: "Birthday Day",
		Required:    true,
	},
}

func SetBirthday(s *discordgo.Session, i *discordgo.InteractionCreate) {

}
