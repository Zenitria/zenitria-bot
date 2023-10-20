package events

import (
	"fmt"
	"math/rand"
	"time"

	"zenitria-bot/database"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.GuildID == "" {
		return
	}

	collection := database.DiscordDB.Collection("Settings")

	var setting database.Setting
	err := collection.FindOne(database.CTX, bson.M{"_id": "Excluded Channels"}).Decode(&setting)

	if err != nil {
		return
	}

	for _, ch := range setting.Value.(primitive.A) {
		if ch == m.ChannelID {
			return
		}
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

	if !manager.CheckUser(m.Author.ID) {
		manager.CreateUser(m.Author.ID)
	}

	user := manager.GetUser(m.Author.ID)

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

	manager.UpdateUser(m.Author.ID, level, xp, nextLevelXP, user.Warnings, cash)

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
