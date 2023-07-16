package economy

import (
	"fmt"
	"zenitria-bot/usermanager"

	"github.com/bwmarrin/discordgo"
)

func HandleBalance(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	var user *discordgo.User

	if len(data.Options) == 0 {
		user = i.Member.User
	} else {
		user = data.Options[0].UserValue(s)
	}

	userInfo := usermanager.GetUser(user.ID)

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🏦・%s's balance", user.Username),
		Description: fmt.Sprintf("💵・**Cash**: %.2f", userInfo.Cash),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.AvatarURL(""),
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