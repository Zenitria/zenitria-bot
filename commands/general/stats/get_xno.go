package stats

import (
	"fmt"
	"zenitria-bot/platforms"

	"github.com/bwmarrin/discordgo"
)

func HandleGetXNO(s *discordgo.Session, i *discordgo.InteractionCreate) {
	platform, change := platforms.GetXNO()

	embed := &discordgo.MessageEmbed{
		Title:       "📊・Get XNO stats",
		Description: fmt.Sprintf("📆・**Days**: %d\n👥・**Users**: %d\n💸・**Withdrawn**: Ӿ%.2f\n📈・**Claims**: %d (%s)", platform.Days, platform.Users, platform.Withdrawn, platform.Claims, change),
		Color:       0xBE4DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://get-xno.com/_next/image?url=%2Fimages%2Fget-xno-logo.png&w=256&q=100",
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
