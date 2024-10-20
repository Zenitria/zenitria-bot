package moderation

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func HandleWarns(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	var user *discordgo.User

	if len(data.Options) == 0 {
		user = i.Member.User
	} else {
		user = data.Options[0].UserValue(s)
	}

	warns := getWarns(user.ID)

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("⚠️・%s's warns", user.Username),
		Description: fmt.Sprintf("⚠️・**Warns**: %d", warns),
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
