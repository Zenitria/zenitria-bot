package settings

import (
	"zenitria-bot/commands/settings/leveling"

	"github.com/bwmarrin/discordgo"
)

func HandleLeveling(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	handlers := map[string](func(*discordgo.Session, *discordgo.InteractionCreate)){
		"exclude":       leveling.HandleExclude,
		"include":       leveling.HandleInclude,
		"excluded-list": leveling.HandleExcludedList,
	}

	if handler, ok := handlers[data.Options[0].Name]; ok {
		handler(s, i)
	}
}
