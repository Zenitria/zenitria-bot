package economy

import (
	"fmt"
	"strings"
	"time"
	"zenitria-bot/code"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleBuy(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckCommandChannel(s, i, i.ChannelID) {
		return
	}

	data := i.ApplicationCommandData()

	item := data.Options[0].StringValue()

	products := map[string]float32{
		"diamonds-mini":    0.5,
		"diamonds-small":   1,
		"diamonds-medium":  3,
		"diamonds-big":     6.5,
		"diamonds-premium": 20,
	}

	if !checkBalance(i.Member.User.ID, products[item]) {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "You don't have enough cash to buy this item.",
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

	if strings.HasPrefix(item, "diamonds-") {
		var diamonds int

		switch item {
		case "diamonds-mini":
			diamonds = 10
		case "diamonds-small":
			diamonds = 25
		case "diamonds-medium":
			diamonds = 100
		case "diamonds-big":
			diamonds = 250
		case "diamonds-premium":
			diamonds = 1000
		}

		c := code.GenerateCode(diamonds, 24, 1)

		updateBalance(i.Member.User.ID, -products[item])

		embed := &discordgo.MessageEmbed{
			Title:       "💎・Diamonds Pack",
			Description: "Successfully bought diamonds pack! Check your DMs for the code",
			Color:       0xB54DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://i.gifer.com/90JG.gif",
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
			Title:       "💎・Diamonds Pack",
			Description: fmt.Sprintf("🏷️・**Code:** %s\n💎・**Diamonds:** %d\n⏳・**Expires:** <t:%d:R>\n💰・**Redeem:** http://get-xno.com/app/redeem", c, diamonds, time.Now().Add(24*time.Hour).Unix()),
			Color:       0xB54DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://i.gifer.com/90JG.gif",
			},
		}

		channel, _ := s.UserChannelCreate(i.Member.User.ID)
		s.ChannelMessageSendEmbed(channel.ID, embed)
	}
}
