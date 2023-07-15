package moderation

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HandleBan(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	user := data.Options[0].UserValue(s)

	var reason string

	if len(data.Options) == 1 {
		reason = "No reason provided."
	} else {
		reason = data.Options[1].StringValue()
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🚷・%s has been banned", user.Username),
		Description: fmt.Sprintf("🚨・**Reason**: %s\n🛡️・**Moderator**: %s", reason, i.Member.User.Mention()),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/TbfChfHKkOUAAAAM/ban-button.gif",
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
		Title:       "🚷・Banned",
		Description: fmt.Sprintf("🚨・**Reason**: %s\n🛡️・**Moderator**: %s", reason, i.Member.User.Mention()),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/TbfChfHKkOUAAAAM/ban-button.gif",
		},
	}

	channel, _ := s.UserChannelCreate(user.ID)
	s.ChannelMessageSendEmbed(channel.ID, embed)

	s.GuildBanCreateWithReason(i.GuildID, user.ID, reason, 0)
}
