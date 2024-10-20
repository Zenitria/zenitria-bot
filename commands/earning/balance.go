package earning

import (
	"fmt"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleBalance(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckCommandChannel(s, i, i.ChannelID) {
		return
	}

	data := i.ApplicationCommandData()

	var user *discordgo.User

	if len(data.Options) == 0 {
		user = i.Member.User
	} else {
		user = data.Options[0].UserValue(s)
	}

	userInfo := manager.GetUser(user.ID)

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🏦・%s's balance", user.Username),
		Description: fmt.Sprintf("💵・**Money**: $%s", formatFloat(userInfo.Cash)),
		Color:       0xBE4DFF,
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
