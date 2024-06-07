package general

import (
	"zenitria-bot/commands/general/stats"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleStats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckCommandChannel(s, i, i.ChannelID) {
		return
	}

	data := i.ApplicationCommandData()

	handlers := map[string](func(*discordgo.Session, *discordgo.InteractionCreate)){
		"get-xno": stats.HandleGetXNO,
		"get-ban": stats.HandleGetBAN,
	}

	if handler, ok := handlers[data.Options[0].Name]; ok {
		handler(s, i)
	}
}
