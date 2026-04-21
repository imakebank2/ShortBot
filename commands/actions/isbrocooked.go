package actions

import (
	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/utils"
)

var responses = []string{"Yes 🥀🙏😭💔", "No 🥶🗿🐺🤫"}

func IsBroCooked(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response, _ := utils.GetRandomElement(responses)
	interactionTextResponse(response, s, i)
}
