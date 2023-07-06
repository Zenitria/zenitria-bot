package usermanager

import (
	"zenitria-bot/database"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckPermissions(perms int64, perm int64, s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	if perms&perm == 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "You don't have permission to use this command.",
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

		return false
	}

	return true
}

func CreateUser(id string) {
	collection := database.DiscordDB.Collection("Users")
	user := database.NewUser(id)

	collection.InsertOne(database.CTX, user)
}

func GetUser(id string) database.User {
	collection := database.DiscordDB.Collection("Users")

	var user database.User

	collection.FindOne(database.CTX, bson.M{"_id": id}).Decode(&user)

	return user
}

func CheckUser(id string) bool {
	collection := database.DiscordDB.Collection("Users")

	var user database.User

	err := collection.FindOne(database.CTX, bson.M{"_id": id}).Decode(&user)

	return err != mongo.ErrNoDocuments
}

func UpdateUser(id string, l int, xp int, n int, w int) {
	collection := database.DiscordDB.Collection("Users")

	update := bson.M{
		"$set": database.User{
			ID:          id,
			Level:       l,
			XP:          xp,
			NextLevelXP: n,
			Warnings:    w,
		},
	}

	collection.UpdateByID(database.CTX, id, update)
}
