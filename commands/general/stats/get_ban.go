package stats

import (
	"fmt"
	"zenitria-bot/platforms"

	"github.com/bwmarrin/discordgo"
)

func HandleGetBAN(s *discordgo.Session, i *discordgo.InteractionCreate) {
	platform, change := platforms.GetBAN()

	embed := &discordgo.MessageEmbed{
		Title:       "📊・Get BAN stats",
		Description: fmt.Sprintf("📆・**Days**: %d\n👥・**Users**: %d\n💸・**Withdrawn**: %.2f BAN\n📈・**Claims**: %d (%s)", platform.Days, platform.Users, platform.Withdrawn, platform.Claims, change),
		Color:       0xBE4DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://get-ban.com/_next/image?url=%2Fimages%2Fget-ban-logo.png&w=256&q=100",
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
