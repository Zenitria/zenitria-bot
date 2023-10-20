package events

import (
	"strings"
	"zenitria-bot/commands/economy"
	"zenitria-bot/commands/general"
	"zenitria-bot/commands/leveling"
	"zenitria-bot/commands/moderation"
	"zenitria-bot/commands/settings"
	"zenitria-bot/manager"

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
			"rank":        leveling.HandleRank,
			"leaderboard": leveling.HandleLeaderboard,
			// Economy
			"balance": economy.HandleBalance,
			"shop":    economy.HandleShop,
			"buy":     economy.HandleBuy,
			"claim":   economy.HandleClaim,
			// Moderation
			"ban":     moderation.HandleBan,
			"unban":   moderation.HandleUnban,
			"kick":    moderation.HandleKick,
			"timeout": moderation.HandleTimeout,
			"warn":    moderation.HandleWarn,
			// Settings
			"set-verification-role":     settings.HandleSetVerificationRole,
			"send-verification-message": settings.HandleSendVerificationMessage,
			"excluded-channels":         settings.HandleExcludedChannels,
			"exclude-channel":           settings.HandleExcludeChannel,
			"include-channel":           settings.HandleIncludeChannel,
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

		handlersWithoutData := map[string](func(*discordgo.Session, *discordgo.InteractionCreate)){
			// Verification
			"verify-button": settings.HandleVerifyButton,
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

		if handler, ok := handlersWithoutData[data.CustomID]; ok {
			handler(s, i)
			return
		}
	case discordgo.InteractionModalSubmit:
		data := i.ModalSubmitData()

		handlers := map[string](func(*discordgo.Session, *discordgo.InteractionCreate, *discordgo.ModalSubmitInteractionData)){
			// Verification
			"verification-modal": manager.HandleVerification,
		}

		if handler, ok := handlers[data.CustomID]; ok {
			handler(s, i, &data)
		}
	}
}
