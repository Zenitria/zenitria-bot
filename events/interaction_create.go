package events

import (
	"strings"
	"zenitria-bot/commands/earning"
	"zenitria-bot/commands/general"
	"zenitria-bot/commands/leveling"
	"zenitria-bot/commands/moderation"
	"zenitria-bot/commands/settings"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		data := i.ApplicationCommandData()

		handlers := map[string](func(*discordgo.Session, *discordgo.InteractionCreate)){
			// General
			"help":  general.HandleHelp,
			"info":  general.HandleInfo,
			"stats": general.HandleStats,
			// Leveling
			"rank":        leveling.HandleRank,
			"leaderboard": leveling.HandleLeaderboard,
			// Earning
			"balance":  earning.HandleBalance,
			"claim":    earning.HandleClaim,
			"withdraw": earning.HandleWithdraw,
			"balances": earning.HandleBalances,
			// Moderation
			"ban":     moderation.HandleBan,
			"unban":   moderation.HandleUnban,
			"kick":    moderation.HandleKick,
			"timeout": moderation.HandleTimeout,
			"warn":    moderation.HandleWarn,
			// Settings
			"leveling": settings.HandleLeveling,
		}

		if handler, ok := handlers[data.Name]; ok {
			handler(s, i)
		}
	case discordgo.InteractionMessageComponent:
		data := i.MessageComponentData()

		handlersWithData := map[string](func(*discordgo.Session, *discordgo.InteractionCreate, string, string)){
			// Help
			"help-menu": general.HandleHelpMenu,
			// Leaderboard
			"leaderboard-previous-button": leveling.HandleLeaderboardButtons,
			"leaderboard-next-button":     leveling.HandleLeaderboardButtons,
		}

		splitted := strings.Split(data.CustomID, "|")

		if handler, ok := handlersWithData[splitted[0]]; ok {
			if len(data.Values) == 0 {
				handler(s, i, splitted[1], string(splitted[2]))
				return
			}

			handler(s, i, splitted[1], data.Values[0])
			return
		}
	}
}
