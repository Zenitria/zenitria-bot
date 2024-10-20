package moderation

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HandleWarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	user := data.Options[0].UserValue(s)

	var reason string

	if len(data.Options) == 1 {
		reason = "No reason provided."
	} else {
		reason = data.Options[1].StringValue()
	}

	addWarn(user.ID)
	warns := getWarns(user.ID)

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("⚠️・%s has been warned", user.Username),
		Description: fmt.Sprintf("🚨・**Reason**: %s\n⚠️・**Warns**: %d\n🛡️・**Moderator**: %s", reason, warns, i.Member.User.Mention()),
		Color:       0xBE4DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/sLgNruA4tsgAAAAC/warning-lights.gif",
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
		Title:       "⚠️・Warn",
		Description: fmt.Sprintf("🚨・**Reason**: %s\n⚠️・**Warns**: %d\n🛡️・**Moderator**: %s", reason, warns, i.Member.User.Mention()),
		Color:       0xBE4DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/sLgNruA4tsgAAAAC/warning-lights.gif",
		},
	}

	channel, _ := s.UserChannelCreate(user.ID)
	s.ChannelMessageSendEmbed(channel.ID, embed)
}
