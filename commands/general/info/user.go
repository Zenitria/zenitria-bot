package info

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HandleUser(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	var user *discordgo.User

	if len(data.Options[0].Options) == 0 {
		user = i.Member.User
	} else {
		user = data.Options[0].Options[0].UserValue(s)
	}

	member, _ := s.GuildMember(i.GuildID, user.ID)

	bot := "🚫"
	createdDate, _ := discordgo.SnowflakeTimestamp(user.ID)

	created := createdDate.Unix()

	nick := member.Nick
	booster := "🚫"
	joined := member.JoinedAt.Unix()
	roles, _ := s.GuildRoles(i.GuildID)
	var role string

	if user.Bot {
		bot = "✅"
	}

	if member.PremiumSince != nil {
		booster = "✅"
	}

	if nick == "" {
		nick = "🚫"
	}

	for _, r := range roles {
		if r.Name == "@everyone" {
			role = r.ID
		}
	}

	if len(member.Roles) > 0 {
		role = member.Roles[0]
	}

	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("👤・Information about %s", user.Username),
		Color: 0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "👤・Account",
				Value: fmt.Sprintf("👤・**Username**: %s\n🆔・**ID**: %s\n🤖・**Bot**: %s\n📆・**Creation Date**: <t:%d>\n\u200b", user.Username, user.ID, bot, created),
			},
			{
				Name:  "🏰・Server",
				Value: fmt.Sprintf("🏷️・**Nick**: %s\n🚀・**Booster**: %s\n📆・**Join Date**: <t:%d>\n🎓・**Top Role**: <@&%s>", nick, booster, joined, role),
			},
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
