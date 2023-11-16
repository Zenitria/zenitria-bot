package leveling

import (
	"fmt"

	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleRank(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	rank, lastRank := getRank(user.ID)

	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🥇・%s's rank", user.Username),
		Description: fmt.Sprintf("✨・**Level**: %d\n🎯・**XP**: %s (%d/%d)\n🥇・**Rank**: %d/%d", userInfo.Level, progressBar(userInfo.XP, userInfo.NextLevelXP), userInfo.XP, userInfo.NextLevelXP, rank, lastRank),
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
