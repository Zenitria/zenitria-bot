package earning

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"zenitria-bot/manager"
)

func checkBalance(id string, am float32) bool {
	user := manager.GetUser(id)

	return user.Cash >= am
}

func updateBalance(id string, am float32) {
	user := manager.GetUser(id)

	user.Cash += am

	manager.UpdateUser(id, user.Level, user.XP, user.NextLevelXP, user.Warnings, user.Cash, user.LastClaimed)
}

func sendInternalError(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "🚫・Error!",
		Description: "Internal error. Please try again later.",
		Color:       0xf66555,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/hI4TN7nt06oAAAAM/error.gif",
		},
	}

	embeds := []*discordgo.MessageEmbed{embed}

	response := &discordgo.WebhookEdit{
		Embeds: &embeds,
	}

	s.InteractionResponseEdit(i.Interaction, response)
}

func sendSuccess(s *discordgo.Session, i *discordgo.InteractionCreate, amount float32, tx, explorer string) {
	embed := &discordgo.MessageEmbed{
		Title:       "💸・Withdraw",
		Description: fmt.Sprintf("The **$%f** was withdrawn successfully!", amount),
		Color:       0xB54DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.gifer.com/90JG.gif",
		},
	}

	embeds := []*discordgo.MessageEmbed{embed}

	response := &discordgo.WebhookEdit{
		Embeds: &embeds,
	}

	s.InteractionResponseEdit(i.Interaction, response)

	embed = &discordgo.MessageEmbed{
		Title:       "💸・Withdraw",
		Description: fmt.Sprintf("💰・**Amount:** $%f\n⛓️・**TxID:** [%s](%s)\n", amount, tx, explorer+tx),
		Color:       0xB54DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.gifer.com/90JG.gif",
		},
	}

	channel, _ := s.UserChannelCreate(i.Member.User.ID)
	s.ChannelMessageSendEmbed(channel.ID, embed)
}
