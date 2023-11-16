package info

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HandleServer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	server, _ := s.State.Guild(i.GuildID)

	created := server.JoinedAt.Unix()

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🌐・Information about %s", server.Name),
		Description: fmt.Sprintf("🏷️・**Name**: %s\n🆔・**ID**: %s\n👑・**Owner**: <@%s>\n👥・**Members**: %d\n🏆・**Boost Tier**: %d\n🚀・**Boosts**: %d\n#️⃣・**Channels**: %d\n🎓・**Roles**: %d\n🙂・**Emojis**: %d\n📆・**Creation Date**: <t:%d>", server.Name, server.ID, server.OwnerID, server.MemberCount, int(server.PremiumTier), server.PremiumSubscriptionCount, len(server.Channels), len(server.Roles), len(server.Emojis), created),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: server.IconURL(""),
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
