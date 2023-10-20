package settings

import (
	"fmt"
	"zenitria-bot/database"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/bwmarrin/discordgo"
)

func HandleSetVerificationRole(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	role := data.Options[0].RoleValue(s, i.GuildID)

	collection := database.DiscordDB.Collection("Settings")

	var setting database.Setting
	err := collection.FindOne(database.CTX, bson.M{"_id": "Verification Role"}).Decode(&setting)

	if err != nil {
		setting = database.Setting{
			Name:  "Verification Role",
			Value: role.ID,
		}

		collection.InsertOne(database.CTX, setting)
	} else {
		collection.UpdateOne(database.CTX, bson.M{"_id": "Verification Role"}, bson.M{"$set": bson.M{"value": role.ID}})
	}

	embed := &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("The verification role has been set to %s.", role.Mention()),
		Color:       0x06e386,
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
