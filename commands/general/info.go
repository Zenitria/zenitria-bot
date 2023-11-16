package general

import (
	"zenitria-bot/commands/general/info"

	"github.com/bwmarrin/discordgo"
)

func HandleInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	handlers := map[string](func(*discordgo.Session, *discordgo.InteractionCreate)){
		"user":   info.HandleUser,
		"server": info.HandleServer,
	}

	if handler, ok := handlers[data.Options[0].Name]; ok {
		handler(s, i)
	}
}
