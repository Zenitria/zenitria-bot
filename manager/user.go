package manager

import (
	"strconv"
	"strings"
	"time"
	"zenitria-bot/database"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func UpdateUser(id string, l int, xp int, n int, w int, c float32) {
	collection := database.DiscordDB.Collection("Users")

	update := bson.M{
		"$set": database.User{
			ID:          id,
			Level:       l,
			XP:          xp,
			NextLevelXP: n,
			Warnings:    w,
			Cash:        c,
		},
	}

	collection.UpdateByID(database.CTX, id, update)
}

func HandleVerification(s *discordgo.Session, i *discordgo.InteractionCreate, data *discordgo.ModalSubmitInteractionData) {
	anwser := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	splitted := strings.Split(data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).CustomID, "|")

	var result int

	first, _ := strconv.Atoi(splitted[0])
	second, _ := strconv.Atoi(splitted[1])

	if splitted[2] == "0" {
		result = first - second
	} else {
		result = first + second
	}

	if anwser != strconv.Itoa(result) {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "You have entered an invalid answer!",
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

	creationTime, _ := discordgo.SnowflakeTimestamp(i.Member.User.ID)
	diff := time.Since(creationTime)
	days := int(diff.Hours() / 24)

	if days < 30 {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "Your account must be at least 30 days old!",
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

	role, _ := GetVerificationRole()

	s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, role)

	embed := &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: "You have been verified!",
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/ikvoQAqXu9MAAAAM/success.gif",
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
}
