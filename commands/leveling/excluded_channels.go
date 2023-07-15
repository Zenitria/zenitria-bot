package leveling

import (
	"github.com/bwmarrin/discordgo"
)

func HandleExcludedChannels(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channels := getExcludedChannels()

	embed := &discordgo.MessageEmbed{
		Title:       "#️⃣・Excluded channels",
		Description: channels,
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/YrtA8A5oLrwAAAAM/excluded.gif",
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
