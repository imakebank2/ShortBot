package birthdays

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/utils"
)

func CheckBirthdays(s *discordgo.Session, path string, channelName string) {
	for {
		loc, err := time.LoadLocation("Australia/Brisbane")

		if err != nil {
			log.Println(err)
			return
		}

		now := time.Now().In(loc)
		nextMidnight := time.Date(
			now.Year(), now.Month(), now.Day()+1,
			0, 0, 0, 0,
			loc,
		)

		waitDuration := time.Until(nextMidnight)
		time.Sleep(waitDuration)

		now = time.Now().In(loc)

		birthdays, err := LoadBirthdays(path)

		if err != nil {
			log.Println(err)
			return
		}

		month := int(now.Month())
		day := now.Day()

		channelID, err := utils.GetChannelIDByName(s, channelName)

		if err != nil {
			log.Println(err)
			return
		}

		for _, b := range birthdays {
			if b.Month == month && b.Day == day {
				message := "@everyone Today is " + b.Name + "'s birthday! 🥳🥳🥳🥳🥳🥳🎉🎉🎉🎉🎉🎉🎉🎉🎉"
				_, err := s.ChannelMessageSend(channelID, message)
				if err != nil {
					log.Println("Error sending message:", err)
				}
			}
		}
	}
}
