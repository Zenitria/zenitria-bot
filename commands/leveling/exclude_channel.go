package leveling

import (
	"fmt"
	"zenitria-bot/usermanager"

	"github.com/bwmarrin/discordgo"
)

func HandleExcludeChannel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	channel := data.Options[0].ChannelValue(s)
	permissions := i.Member.Permissions

	if !usermanager.CheckPermissions(permissions, discordgo.PermissionManageChannels, s, i) {
		return
	}

	if !checkChannelInDB(channel.ID) {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: fmt.Sprintf("The channel %s is already excluded from leveling system.", channel.Mention()),
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

		return
	}

	addChannelToDB(channel.ID)

	embed := &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("The channel %s has been excluded from leveling system.", channel.Mention()),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/ikvoQAqXu9MAAAAM/success.gif",
		},
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	}

	s.InteractionRespond(i.Interaction, response)
}
