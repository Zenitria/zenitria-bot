package general

import (
	"fmt"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckCommandChannel(s, i, i.ChannelID) {
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "📚・Help",
		Description: "Select a help category to see more information about the commands.",
		Color:       0xBE4DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
		},
	}

	selectMenu := &discordgo.SelectMenu{
		CustomID:    fmt.Sprintf("help-menu|%s", i.Member.User.ID),
		Placeholder: "Select a category",
		Options: []discordgo.SelectMenuOption{
			{
				Label: "General",
				Value: "general",
				Emoji: discordgo.ComponentEmoji{
					Name: "📖",
				},
			},
			{
				Label: "Leveling",
				Value: "leveling",
				Emoji: discordgo.ComponentEmoji{
					Name: "✨",
				},
			},
			{
				Label: "Earning",
				Value: "earning",
				Emoji: discordgo.ComponentEmoji{
					Name: "💵",
				},
			},
			{
				Label: "Moderation",
				Value: "moderation",
				Emoji: discordgo.ComponentEmoji{
					Name: "🛡️",
				},
			},
			{
				Label: "Settings",
				Value: "settings",
				Emoji: discordgo.ComponentEmoji{
					Name: "⚙️",
				},
			},
		},
	}

	components := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{selectMenu},
		},
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: components,
		},
	}

	s.InteractionRespond(i.Interaction, response)
}

func HandleHelpMenu(s *discordgo.Session, i *discordgo.InteractionCreate, id string, val string) {
	if id != i.Member.User.ID {
		return
	}

	selectMenu := &discordgo.SelectMenu{
		CustomID:    fmt.Sprintf("help-menu|%s", i.Member.User.ID),
		Placeholder: "Select a category",
		Options: []discordgo.SelectMenuOption{
			{
				Label: "General",
				Value: "general",
				Emoji: discordgo.ComponentEmoji{
					Name: "📖",
				},
			},
			{
				Label: "Leveling",
				Value: "leveling",
				Emoji: discordgo.ComponentEmoji{
					Name: "✨",
				},
			},
			{
				Label: "Earning",
				Value: "earning",
				Emoji: discordgo.ComponentEmoji{
					Name: "💵",
				},
			},
			{
				Label: "Moderation",
				Value: "moderation",
				Emoji: discordgo.ComponentEmoji{
					Name: "🛡️",
				},
			},
			{
				Label: "Settings",
				Value: "settings",
				Emoji: discordgo.ComponentEmoji{
					Name: "⚙️",
				},
			},
		},
	}

	components := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{selectMenu},
		},
	}

	switch val {
	case "general":
		mentions := getSlashCommandMentions(s)
		submmentions := getSlashSubcommandMentions(s)

		embed := &discordgo.MessageEmbed{
			Title: "📖・General",
			Color: 0xBE4DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  submmentions["info user"] + " (user)",
					Value: "Shows information about yourself or another user.",
				},
				{
					Name:  submmentions["info server"],
					Value: "Shows information about the server.",
				},
				{
					Name:  submmentions["stats get-xno"],
					Value: "Shows the stats of Get XNO.",
				},
				{
					Name:  submmentions["stats get-ban"],
					Value: "Shows the stats of Get BAN.",
				},
				{
					Name:  mentions["send-code"],
					Value: "Sends code to the codes channel.",
				},
				{
					Name:  mentions["help"],
					Value: "Shows this help menu.",
				},
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
			},
		}

		s.InteractionRespond(i.Interaction, response)
	case "leveling":
		mentions := getSlashCommandMentions(s)

		embed := &discordgo.MessageEmbed{
			Title: "✨・Leveling",
			Color: 0xBE4DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  mentions["rank"] + " (user)",
					Value: "Shows your rank or the rank of another user.",
				},
				{
					Name:  mentions["leaderboard"],
					Value: "Shows the server's leaderboard.",
				},
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
			},
		}

		s.InteractionRespond(i.Interaction, response)
	case "earning":
		mentions := getSlashCommandMentions(s)

		embed := &discordgo.MessageEmbed{
			Title: "💵・Earning",
			Color: 0xBE4DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  mentions["balance"] + " (user)",
					Value: "Shows your balance or the balance of another user.",
				},
				{
					Name:  mentions["withdraw"],
					Value: "Withdraws your balance.",
				},
				{
					Name:  mentions["claim"],
					Value: "Claims your 15 minute reward.",
				},
				{
					Name:  mentions["balances"],
					Value: "Shows bot crypto balances.",
				},
				{
					Name:  mentions["rich-list"],
					Value: "Shows the server's rich list.",
				},
				{
					Name:  mentions["add-dollars"] + " [user] [amount]",
					Value: "Adds dollars to a user.",
				},
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
			},
		}

		s.InteractionRespond(i.Interaction, response)
	case "moderation":
		mentions := getSlashCommandMentions(s)

		embed := &discordgo.MessageEmbed{
			Title: "🛡️・Moderation",
			Color: 0xBE4DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  mentions["ban"] + " [user] (reason)",
					Value: "Bans a user from the server.",
				},
				{
					Name:  mentions["unban"] + " [user id]",
					Value: "Unbans a user from the server.",
				},
				{
					Name:  mentions["kick"] + " [user] (reason)",
					Value: "Kicks a user from the server.",
				},
				{
					Name:  mentions["timeout"] + " [user] [duration] (reason)",
					Value: "Timeout a user.",
				},
				{
					Name:  mentions["warn"] + " [user] (reason)",
					Value: "Warns a user.",
				},
				{
					Name:  mentions["warns"] + " (user)",
					Value: "Shows your warns or the warns of another user.",
				},
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
			},
		}

		s.InteractionRespond(i.Interaction, response)
	case "settings":
		mentions := getSlashSubcommandMentions(s)

		embed := &discordgo.MessageEmbed{
			Title: "⚙️・Settings",
			Color: 0xBE4DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  mentions["leveling excluded-list"],
					Value: "Lists all excluded channels from the leveling system.",
				},
				{
					Name:  mentions["leveling exclude"] + " [channel]",
					Value: "Excludes a channel from the leveling system.",
				},
				{
					Name:  mentions["leveling include"] + " [channel]",
					Value: "Includes a channel in the leveling system.",
				},
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
			},
		}

		s.InteractionRespond(i.Interaction, response)
	}
}
