package events

import (
	"strings"
	"zenitria-bot/commands/general"
	"zenitria-bot/commands/leveling"
	"zenitria-bot/commands/moderation"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()

		handlers := map[string](func(*discordgo.Session, *discordgo.InteractionCreate)){
			// General
			"user-info":   general.HandleUserInfo,
			"server-info": general.HandleServerInfo,
			"get-xno":     general.HandleGetXNO,
			"help":        general.HandleHelp,
			// Leveling
			"rank":              leveling.HandleRank,
			"leaderboard":       leveling.HandleLeaderboard,
			"excluded-channels": leveling.HandleExcludedChannels,
			"exclude-channel":   leveling.HandleExcludeChannel,
			"include-channel":   leveling.HandleIncludeChannel,
			// Moderation
			"ban":     moderation.HandleBan,
			"unban":   moderation.HandleUnban,
			"kick":    moderation.HandleKick,
			"timeout": moderation.HandleTimeout,
		}

		if handler, ok := handlers[data.Name]; ok {
			handler(s, i)
		}
	case discordgo.InteractionMessageComponent:
		data := i.MessageComponentData()

		handlers := map[string](func(*discordgo.Session, *discordgo.InteractionCreate, string, string)){
			"help-menu":                   general.HandleHelpMenu,
			"leaderboard-previous-button": leveling.HandleLeaderboardButtons,
			"leaderboard-next-button":     leveling.HandleLeaderboardButtons,
		}

		splitted := strings.Split(data.CustomID, "|")

		if handler, ok := handlers[splitted[0]]; ok {
			if len(data.Values) == 0 {
				handler(s, i, splitted[1], string(splitted[2]))
				return
			}

			handler(s, i, splitted[1], data.Values[0])
		}
	}
}
