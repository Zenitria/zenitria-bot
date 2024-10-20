package earning

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"zenitria-bot/database"
	"zenitria-bot/manager"
)

func HandleAddDollars(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	member := data.Options[0].UserValue(s)
	amount := data.Options[1].FloatValue()

	if amount < 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "You can't add negative dollars!",
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

	collection := database.DiscordDB.Collection("Users")

	if !manager.CheckUser(member.ID) {
		manager.CreateUser(member.ID)
	}

	var user database.User
	collection.FindOne(database.CTX, bson.M{"_id": member.ID}).Decode(&user)

	user.Cash += amount
	manager.UpdateUser(member.ID, user.Level, user.XP, user.NextLevelXP, user.Warnings, user.Cash, user.LastClaimed)

	embed := &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("Successfully added $%s to %s's balance!", formatFloat(amount), member.Mention()),
		Color:       0xBE4DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/ikvoQAqXu9MAAAAM/success.gif",
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
