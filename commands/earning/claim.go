package earning

import (
	"fmt"
	"time"
	"zenitria-bot/database"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleClaim(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckCommandChannel(s, i, i.ChannelID) {
		return
	}

	collection := database.DiscordDB.Collection("Users")

	if !manager.CheckUser(i.Member.User.ID) {
		manager.CreateUser(i.Member.User.ID)
	}

	var user database.User
	collection.FindOne(database.CTX, bson.M{"_id": i.Member.User.ID}).Decode(&user)

	nextClaimDate := user.LastClaimed.Add(15 * time.Minute)

	if nextClaimDate.After(time.Now()) {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: fmt.Sprintf("You have already claimed your 15 minute reward! Next claim <t:%d:R>.", nextClaimDate.Unix()),
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

	embed := &discordgo.MessageEmbed{
		Title:       "💰・Claim",
		Description: fmt.Sprintf("[Click here](https://zenitria.dev/bot/claim/%s) to claim your 15 minute reward!", i.Member.User.ID),
		Color:       0xBE4DFF,
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
