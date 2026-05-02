package birthdays

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/utils"
)

func CheckBirthdays(s *discordgo.Session, path string, channelName string) {
	loc, err := time.LoadLocation("Australia/Brisbane")
	if err != nil {
		log.Println(err)
		return
	}

	// Check if birthdays can actually be loaded
	birthdays, err := LoadBirthdays(path)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Birthdays:", fmt.Sprintf("%v", birthdays))
	}

	channelID, err := utils.GetChannelIDByName(s, channelName)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		now := time.Now().In(loc)
		nextDay := time.Date(
			now.Year(), now.Month(), now.Day(),
			0, 0, 1, 0,
			loc,
		).AddDate(0, 0, 1)

		waitDuration := time.Until(nextDay)
		time.Sleep(waitDuration)

		now = time.Now().In(loc)

		birthdays, err := LoadBirthdays(path)

		if err != nil {
			log.Println(err)
			return
		}

		month := int(now.Month())
		day := now.Day()

		for _, b := range birthdays {
			if b.Month == month && b.Day == day {
				message := "@everyone Today is " + b.Name + "'s birthday! Happy Birthday! 🥳🎉"
				_, err := s.ChannelMessageSend(channelID, message)
				if err != nil {
					log.Println("Error sending message:", err)
				}
			}
		}
	}
}
