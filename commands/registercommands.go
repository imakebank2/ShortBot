package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/commands/actions"
)

// Add your commands here
var BotCommands = []Command{
	{"quoteoftheday", "Says a random quote from #shortgang-quotes", actions.QuoteOfTheDay, nil},
	{"randomuser", "Gets a random user", actions.RandomUser, nil},
	{"67", "Info about the 67.", actions.SixEvenInfo, nil},
	{"1v1", "1v1 someone", actions.OneVsOne, actions.OneVsOneOptions},
	{"isbrocooked", "Is brotato chip cooked to the brim 🥀😭", actions.IsBroCooked, nil},
	{"tungtungtungsahur", "Do not try this at home.", actions.TungTungTungSahur, nil},
	{"birthdays", "Get all birthdays", actions.GetBirthdays, nil},
}

func RegisterCommands(s *discordgo.Session) {
	// Deregister existing commands
	// u, err := s.User("@me")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// appID := u.ID

	// commands, _ := s.ApplicationCommands(appID, "")

	// for _, cmd := range commands {
	// 	err := s.ApplicationCommandDelete(appID, "", cmd.ID)
	// 	if err != nil {
	// 		log.Println("failed to delete:", cmd.Name, err)
	// 	} else {
	// 		log.Println("deleted:", cmd.Name)
	// 	}
	// }

	for _, c := range BotCommands {
		c.Register(s)
	}
}
