package general

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HandleHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "📚・Help",
		Description: "Select a help category to see more information about the commands.",
		Color:       0x06e386,
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
				Label: "Economy",
				Value: "economy",
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
				Label: "Economy",
				Value: "economy",
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

		embed := &discordgo.MessageEmbed{
			Title: "📖・General",
			Color: 0x06e386,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  mentions["user-info"] + " (user)",
					Value: "Shows information about yourself or another user.",
				},
				{
					Name:  mentions["server-info"],
					Value: "Shows information about the server.",
				},
				{
					Name:  mentions["get-xno"],
					Value: "Shows the stats of Get XNO.",
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
			Color: 0x06e386,
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
	case "economy":
		mentions := getSlashCommandMentions(s)

		embed := &discordgo.MessageEmbed{
			Title: "💵・Economy",
			Color: 0x06e386,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  mentions["balance"] + " (user)",
					Value: "Shows your balance or the balance of another user.",
				},
				{
					Name:  mentions["shop"],
					Value: "Shows the shop.",
				},
				{
					Name:  mentions["buy"] + " [item]",
					Value: "Buys an item from the shop.",
				},
				{
					Name:  mentions["claim"],
					Value: "Claims your hourly reward.",
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
			Color: 0x06e386,
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
			Color: 0x06e386,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gifdb.com/images/high/need-help-baby-in-lamp-22l1cd6hndd62nkl.gif",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  mentions["verification role"] + " [role]",
					Value: "Sets the verification role.",
				},
				{
					Name:  mentions["verification send"] + " [channel]",
					Value: "Sends the verification message to selected channel",
				},
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
