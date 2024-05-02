package leveling

import (
	"strconv"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleLeaderboard(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckCommandChannel(s, i, i.ChannelID) {
		return
	}

	page := 1

	embed, components := createLeaderboardEmbedAndComponents(i, page)

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: components,
		},
	}

	s.InteractionRespond(i.Interaction, response)
}

func HandleLeaderboardButtons(s *discordgo.Session, i *discordgo.InteractionCreate, id string, p string) {
	if id != i.Member.User.ID {
		return
	}

	page, _ := strconv.Atoi(p)

	embed, components := createLeaderboardEmbedAndComponents(i, page)

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: components,
		},
	}

	s.InteractionRespond(i.Interaction, response)
}
