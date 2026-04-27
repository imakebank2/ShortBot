package utils

import (
	"fmt"
	"math/rand/v2"

	"github.com/bwmarrin/discordgo"
)

func GetRandomElement[T any](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, fmt.Errorf("Slice is empty.")
	}

	return s[rand.IntN(len(s))], nil
}

func GetChannelIDByName(s *discordgo.Session, channelName string) (string, error) {
	for _, guild := range s.State.Guilds {

		channels, err := s.GuildChannels(guild.ID)
		if err != nil {
			return "", err
		}

		for _, c := range channels {
			if c.Name == channelName {
				return c.ID, nil
			}
		}
	}

	return "", fmt.Errorf("Channel ID not found")
}
