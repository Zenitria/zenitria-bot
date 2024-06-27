package general

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"zenitria-bot/codes"
	"zenitria-bot/config"
)

func HandleSendCode(s *discordgo.Session, i *discordgo.InteractionCreate) {
	code, expires := codes.GenerateCode(50, 24, 0)

	embed := &discordgo.MessageEmbed{
		Title:       "💎・Diamonds Code",
		Description: fmt.Sprintf("🏷️・**Code:** %s\n💎・**Diamonds:** 50\n⏳・**Expires:** <t:%d:R>\n💰・**Redeem:** [Get XNO](https://get-xno.com/app/redeem) & [Get BAN](https://get-ban.com/app/redeem)", code, expires.Unix()),
		Color:       0xB54DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/SAJ5PrWD0DcAAAAC/diamond.gif",
		},
	}

	s.ChannelMessageSendEmbed(config.CODES_CHANNEL_ID, embed)
	s.ChannelMessageSend(config.CODES_CHANNEL_ID, fmt.Sprintf("<@&%s>", config.CODES_ROLE_ID))

	embed = &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("The code has been sent to <#%s> channel.", config.CODES_CHANNEL_ID),
		Color:       0xB54DFF,
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
