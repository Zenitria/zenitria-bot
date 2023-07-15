package events

import (
	"fmt"
	"math/rand"
	"time"

	"zenitria-bot/database"
	"zenitria-bot/usermanager"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.GuildID == "" {
		return
	}

	collection := database.DiscordDB.Collection("Excluded Channels")

	err := collection.FindOne(database.CTX, bson.M{"_id": m.ChannelID}).Err()

	if err == nil {
		return
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var randNum int

	if len(m.Content) <= 120 {
		randNum = rng.Intn(2) + 1
	} else if len(m.Content) <= 300 {
		randNum = rng.Intn(3) + 2
	} else {
		randNum = rng.Intn(3) + 4
	}

	if !usermanager.CheckUser(m.Author.ID) {
		usermanager.CreateUser(m.Author.ID)
	}

	user := usermanager.GetUser(m.Author.ID)

	level := user.Level
	xp := user.XP + randNum
	nextLevelXP := user.NextLevelXP
	levelUP := false
	cash := user.Cash + (float32(randNum) / 100)

	if xp >= nextLevelXP {
		xp -= nextLevelXP
		level++
		nextLevelXP = int(float64(nextLevelXP) * 1.5)
		levelUP = true
	}

	usermanager.UpdateUser(m.Author.ID, level, xp, nextLevelXP, user.Warnings, cash)

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
}
