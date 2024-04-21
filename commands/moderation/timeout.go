package moderation

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func HandleTimeout(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	user := data.Options[0].UserValue(s)
	duration := data.Options[1].IntValue()
	durationTime := time.Now().Add(time.Duration(duration) * time.Second)

	durationString := getDurationString(duration)

	var reason string

	if len(data.Options) == 2 {
		reason = "No reason provided."
	} else {
		reason = data.Options[2].StringValue()
	}

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🔇・%s has been timed out", user.Username),
		Description: fmt.Sprintf("🚨・**Reason**: %s\n⏳・**Duration**: %s\n🛡️・**Moderator**: %s", reason, durationString, i.Member.User.Mention()),
		Color:       0xB54DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media1.giphy.com/media/A9FvmJdp3F8hNZK9Ra/200w.gif",
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
		Title:       "🔇・Timeout",
		Description: fmt.Sprintf("🚨・**Reason**: %s\n⏳・**Duration**: %s\n🛡️・**Moderator**: %s", reason, durationString, i.Member.User.Mention()),
		Color:       0xB54DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media1.giphy.com/media/A9FvmJdp3F8hNZK9Ra/200w.gif",
		},
	}

	channel, _ := s.UserChannelCreate(user.ID)
	s.ChannelMessageSendEmbed(channel.ID, embed)

	s.GuildMemberTimeout(i.GuildID, user.ID, &durationTime)
}
