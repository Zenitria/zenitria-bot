package moderation

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HandleBan(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	user := data.Options[0].UserValue(s)

	var reason string
	var deleteDays int

	if len(data.Options) == 1 {
		reason = "No reason provided."
	} else if len(data.Options) == 2 && data.Options[1].Type == discordgo.ApplicationCommandOptionString {
		reason = data.Options[1].StringValue()
	} else if len(data.Options) == 2 && data.Options[1].Type == discordgo.ApplicationCommandOptionInteger {
		reason = "No reason provided."
		deleteDays = int(data.Options[1].IntValue())
	} else {
		reason = data.Options[1].StringValue()
		deleteDays = int(data.Options[2].IntValue())
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

	s.GuildBanCreateWithReason(i.GuildID, user.ID, reason, deleteDays)
}
