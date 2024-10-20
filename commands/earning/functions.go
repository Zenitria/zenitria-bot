package earning

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"strings"
	"zenitria-bot/database"
	"zenitria-bot/manager"
)

func checkBalance(id string, am float64) bool {
	user := manager.GetUser(id)

	return user.Cash >= am
}

func updateBalance(id string, am float64) {
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

func sendSuccess(s *discordgo.Session, i *discordgo.InteractionCreate, am float64, tx, explorer string) {
	embed := &discordgo.MessageEmbed{
		Title:       "💸・Withdraw",
		Description: fmt.Sprintf("The **$%f** was withdrawn successfully!", am),
		Color:       0xBE4DFF,
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
		Description: fmt.Sprintf("💰・**Amount:** $%f\n⛓️・**TxID:** [%s](%s)\n", am, tx, explorer+tx),
		Color:       0xBE4DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.gifer.com/90JG.gif",
		},
	}

	channel, _ := s.UserChannelCreate(i.Member.User.ID)
	s.ChannelMessageSendEmbed(channel.ID, embed)
}

func getRichListUsers() []database.User {
	collection := database.DiscordDB.Collection("Users")

	var users []database.User

	options := options.Find().SetSort(bson.D{
		{Key: "cash", Value: -1},
	})

	cursor, err := collection.Find(database.CTX, bson.M{}, options)

	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(database.CTX) {
		var user database.User
		cursor.Decode(&user)

		users = append(users, user)
	}

	return users
}

func getRichList(p int) string {
	users := getRichListUsers()

	var output string

	for i := (p - 1) * 10; i < p*10; i++ {
		if i >= len(users) {
			break
		}

		output += fmt.Sprintf("%d. <@%s> - $%f\n", i+1, users[i].ID, users[i].Cash)
	}

	if output == "" {
		output = "No users in rich list."
	}

	return output
}

func getRichListPages() int {
	users := getRichListUsers()
	pages := int(math.Ceil(float64(len(users)) / 10))

	if pages == 0 {
		pages = 1
	}

	return pages
}

func createRichListEmbedAndComponents(i *discordgo.InteractionCreate, p int) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	richList := getRichList(p)
	pages := getRichListPages()

	embed := &discordgo.MessageEmbed{
		Title:       "💰・Rich List",
		Color:       0xBE4DFF,
		Description: richList,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Page %d/%d", p, pages),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/LOu06WCOACcAAAAM/money-rich.gif",
		},
	}

	prevButton := &discordgo.Button{
		Emoji: discordgo.ComponentEmoji{
			Name: "◀️",
		},
		Style:    discordgo.PrimaryButton,
		Disabled: p == 1,
		CustomID: fmt.Sprintf("rich-list-previous-button|%s|%d", i.Member.User.ID, p-1),
	}

	nextButton := &discordgo.Button{
		Emoji: discordgo.ComponentEmoji{
			Name: "▶️",
		},
		Style:    discordgo.PrimaryButton,
		Disabled: p == pages,
		CustomID: fmt.Sprintf("rich-list-next-button|%s|%d", i.Member.User.ID, p+1),
	}

	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				prevButton,
				nextButton,
			},
		},
	}

	return embed, components
}

func formatFloat(f float64) string {
	s := fmt.Sprintf("%.15f", f)
	s = strings.TrimRight(s, "0")

	if s[len(s)-1] == '.' {
		s += "00"
	}

	return s
}
