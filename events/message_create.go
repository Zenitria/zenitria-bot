package events

import (
	"fmt"
	"math/rand"
	"time"

	"zenitria-bot/code"
	"zenitria-bot/usermanager"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.GuildID == "" {
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var randomXP int

	if len(m.Content) <= 120 {
		randomXP = rng.Intn(2) + 1
	} else if len(m.Content) <= 300 {
		randomXP = rng.Intn(3) + 2
	} else {
		randomXP = rng.Intn(3) + 4
	}

	if !usermanager.CheckUser(m.Author.ID) {
		usermanager.CreateUser(m.Author.ID)
	}

	user := usermanager.GetUser(m.Author.ID)

	level := user.Level
	xp := user.XP + randomXP
	nextLevelXP := user.NextLevelXP
	levelUP := false

	if xp >= nextLevelXP {
		xp -= nextLevelXP
		level++
		nextLevelXP = int(float64(nextLevelXP) * 1.5)
		levelUP = true
	}

	usermanager.UpdateUser(m.Author.ID, level, xp, nextLevelXP, user.Warnings)

	if !levelUP {
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "✨・Level UP",
		Description: fmt.Sprintf("Congratulations! %s gained level %d.", m.Author.Mention(), level),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/Duc7gUlXkdYAAAAC/level-up.gif",
		},
	}

	s.ChannelMessageSendEmbedReply(m.ChannelID, embed, m.Reference())

	code := code.GenerateCode(level*10, 24, 1)

	embed = &discordgo.MessageEmbed{
		Title:       "✨・Level UP",
		Description: "Congratulations on your new level!",
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/Duc7gUlXkdYAAAAC/level-up.gif",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Reward",
				Value: fmt.Sprintf("🎁・**Code**: %s\n💎・**Diamonds**: %d\n⏳・**Expires**: <t:%d:R>\n💰・**Redeem**: http://get-xno.com/app/redeem", code, level*10, time.Now().Add(24*time.Hour).Unix()),
			},
		},
	}

	channel, _ := s.UserChannelCreate(m.Author.ID)
	s.ChannelMessageSendEmbed(channel.ID, embed)
}
