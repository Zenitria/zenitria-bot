package earning

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"zenitria-bot/manager"
)

func HandleRichList(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckCommandChannel(s, i, i.ChannelID) {
		return
	}

	page := 1
	embed, components := createRichListEmbedAndComponents(i, page)

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: components,
		},
	}

	s.InteractionRespond(i.Interaction, response)
}

func HandleRichListButtons(s *discordgo.Session, i *discordgo.InteractionCreate, id string, p string) {
	if id != i.Member.User.ID {
		return
	}

	page, _ := strconv.Atoi(p)
	embed, components := createRichListEmbedAndComponents(i, page)

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: components,
		},
	}

	s.InteractionRespond(i.Interaction, response)
}
