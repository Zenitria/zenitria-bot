package leveling

import (
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleExcludedList(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channels := manager.GetExcludedChannels()

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
