package leveling

import (
	"fmt"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleExclude(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	channel := data.Options[0].Options[0].ChannelValue(s)

	if manager.IsChannelExcluded(channel.ID) {
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

	manager.ExcludeChannel(channel.ID)

	embed := &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("The channel %s has been excluded from leveling system.", channel.Mention()),
		Color:       0xBE4DFF,
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
