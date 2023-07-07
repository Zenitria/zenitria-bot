package commands

import "github.com/bwmarrin/discordgo"

func GuildChecker(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	if i.GuildID == "" {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "This command can only be used in a server.",
			Color:       0xf66555,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://media.tenor.com/hI4TN7nt06oAAAAM/error.gif",
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
				Flags:  discordgo.MessageFlagsEphemeral,
			},
		}

		s.InteractionRespond(i.Interaction, response)
		return false
	}

	return true
}
