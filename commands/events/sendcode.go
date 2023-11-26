package events

import (
	"fmt"
	"time"
	"zenitria-bot/code"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func SendCodeHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckOwner(s, i) {
		return
	}

	data := i.ApplicationCommandData()
	user := data.Options[0].UserValue(s)
	diamonds := data.Options[1].IntValue()

	c := code.GenerateCode(int(diamonds), 24, 1)

	embed := &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("Successfully sent a code to <@%s>.", user.ID),
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

	embed = &discordgo.MessageEmbed{
		Title:       "💎・Diamonds",
		Description: fmt.Sprintf("🏷️・**Code:** %s\n💎・**Diamonds:** %d\n⏳・**Expires:** <t:%d:R>\n💰・**Redeem:** http://get-xno.com/app/redeem", c, diamonds, time.Now().Add(24*time.Hour).Unix()),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/SAJ5PrWD0DcAAAAC/diamond.gif",
		},
	}

	channel, _ := s.UserChannelCreate(user.ID)
	s.ChannelMessageSendEmbed(channel.ID, embed)
}
