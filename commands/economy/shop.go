package economy

import "github.com/bwmarrin/discordgo"

func HandleShop(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "🛒・Shop",
		Description: "You can buy items with your cash here!",
		Color:       0x06e386,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "💎・Diamonds Packs",
				Value: "✨・**Mini (10):** 💵 0.50\n🌟・**Small (25):** 💵 1.00\n🎁・**Medium (100):** 💵 3.00\n🔥・**Big (250):** 💵 6.50\n🏆・**Premium (1000):** 💵 20.00",
			},
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
