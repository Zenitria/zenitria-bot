package manager

import (
	"fmt"
	"zenitria-bot/config"

	"github.com/bwmarrin/discordgo"
)

func CheckCommandChannel(s *discordgo.Session, i *discordgo.InteractionCreate, id string) bool {
	if id != config.COMMANDS_CHANNEL_ID {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: fmt.Sprintf("You can only use commands on <#%s> channel.", config.COMMANDS_CHANNEL_ID),
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
	}

	return id != config.COMMANDS_CHANNEL_ID
}

func CheckOwner(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	if i.Member.User.ID != config.OWNER_ID {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "Only the owner of the server can use this command.",
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
	}

	return i.Member.User.ID != config.OWNER_ID
}
