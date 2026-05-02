package actions

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/birthdays"
)

func GetBirthdays(s *discordgo.Session, i *discordgo.InteractionCreate) {
	bdays, err := birthdays.LoadBirthdays("./birthdays/birthdays.json")

	if err != nil {
		log.Println(err)
		return
	}

	var reply string

	for _, b := range bdays {
		reply += fmt.Sprintf("%s: %02d/%02d\n", b.Name, b.Day, b.Month)
	}

	interactionTextResponse(reply+"\nIf your birthday isn't on here or is incorrect tell me.", s, i)
}
