package economy

import (
	"fmt"
	"time"
	"zenitria-bot/database"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleClaim(s *discordgo.Session, i *discordgo.InteractionCreate) {
	collection := database.DiscordDB.Collection("Users")

	if !manager.CheckUser(i.Member.User.ID) {
		manager.CreateUser(i.Member.User.ID)
	}

	var user database.User
	collection.FindOne(database.CTX, bson.M{"_id": i.Member.User.ID}).Decode(&user)

	nextClaimDate := user.LastClaimed.Add(1 * time.Hour)

	if nextClaimDate.After(time.Now()) {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: fmt.Sprintf("You have already claimed your hourly reward! Next claim <t:%d:R>.", nextClaimDate.Unix()),
			Color:       0xf66555,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://media.tenor.com/hI4TN7nt06oAAAAM/error.gif",
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
				Flags:  discordgo.MessageFlagsEphemeral,
			},
		}

		s.InteractionRespond(i.Interaction, response)
		return
	}

	amount := 0.1 + (float32(user.Level) * 0.1)

	manager.UpdateUser(i.Member.User.ID, user.Level, user.XP, user.NextLevelXP, user.Warnings, user.Cash+amount, time.Now())

	embed := &discordgo.MessageEmbed{
		Title:       "💰・Claim",
		Description: fmt.Sprintf("You have received 💵 **%.2f**.", amount),
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/6Hixx4SFAeQAAAAM/backing-you-get-yours.gif",
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
