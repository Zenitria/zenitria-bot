package leveling

import (
	"fmt"
	"math"
	"zenitria-bot/database"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func progressBar(xp int, nextLevelXP int) string {
	var output string

	if nextLevelXP == 0 {
		return "░░░░░░░░░░"
	}

	for i := 1; i <= 10; i++ {
		if xp >= nextLevelXP/10*i {
			output += "█"
		} else {
			output += "░"
		}
	}

	return output
}

func getLeaderboardUsers() []database.User {
	collection := database.DiscordDB.Collection("Users")

	var users []database.User

	options := options.Find().SetSort(bson.D{
		{Key: "level", Value: -1},
		{Key: "xp", Value: -1},
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

func getRank(id string) (int, int) {
	leaderboard := getLeaderboardUsers()

	for i, u := range leaderboard {
		if u.ID == id {
			return i + 1, len(leaderboard)
		}
	}

	return 0, len(leaderboard)
}

func getLeaderboard(p int) string {
	users := getLeaderboardUsers()

	var output string

	for i := (p - 1) * 10; i < p*10; i++ {
		if i >= len(users) {
			break
		}

		output += fmt.Sprintf("%d. <@%s> - Level %d (%d XP)\n", i+1, users[i].ID, users[i].Level, users[i].XP)
	}

	if output == "" {
		output = "No users in leaderboard."
	}

	return output
}

func getLeaderboardPages() int {
	users := getLeaderboardUsers()
	pages := int(math.Ceil(float64(len(users)) / 10))

	if pages == 0 {
		pages = 1
	}

	return pages
}

func createLeaderboardEmbedAndComponents(i *discordgo.InteractionCreate, p int) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	leaderboard := getLeaderboard(p)

	pages := getLeaderboardPages()

	embed := &discordgo.MessageEmbed{
		Title:       "🏆・Leaderboard",
		Color:       0xBE4DFF,
		Description: leaderboard,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Page %d/%d", p, pages),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExNHltZ2dqand6MHVyOXl5MzUwN244azkzYnBxN2R0dXd5Z2k3dGN5bCZlcD12MV9naWZzX3NlYXJjaCZjdD1n/vIJaz7nMJhTUc/giphy.gif",
		},
	}

	prevButton := &discordgo.Button{
		Emoji: discordgo.ComponentEmoji{
			Name: "◀️",
		},
		Style:    discordgo.PrimaryButton,
		Disabled: p == 1,
		CustomID: fmt.Sprintf("leaderboard-previous-button|%s|%d", i.Member.User.ID, p-1),
	}

	nextButton := &discordgo.Button{
		Emoji: discordgo.ComponentEmoji{
			Name: "▶️",
		},
		Style:    discordgo.PrimaryButton,
		Disabled: p == pages,
		CustomID: fmt.Sprintf("leaderboard-next-button|%s|%d", i.Member.User.ID, p+1),
	}

	components := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				prevButton,
				nextButton,
			},
		},
	}

	return embed, components
}
