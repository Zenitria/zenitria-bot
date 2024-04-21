package events

import (
	"fmt"
	"time"
	"zenitria-bot/code"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

var (
	codes = []string{}
	setup = false
)

func HandleAdventCalendar(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckOwner(s, i) {
		return
	}

	if setup {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "The advent calendar is already set.",
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

	data := i.ApplicationCommandData()
	channel := data.Options[0].ChannelValue(s)

	fields := []*discordgo.MessageEmbedField{}

	for i := 0; i < 24; i++ {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("🎁・Day %d", i+1),
			Value:  "||#######||",
			Inline: true,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:  "🎄・Advent Calendar",
		Fields: fields,
		Color:  0xB54DFF,
	}

	msg, _ := s.ChannelMessageSendEmbed(channel.ID, embed)

	embed = &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("The advent calendar has been set on %s channel.", channel.Mention()),
		Color:       0xB54DFF,
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

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Cron("0 0 * * *").LimitRunsTo(24).Do(func() {
		cron(s, i, channel, msg)
	})

	scheduler.StartAsync()

	setup = true
}

func cron(s *discordgo.Session, i *discordgo.InteractionCreate, channel *discordgo.Channel, msg *discordgo.Message) {
	now := time.Now().UTC()

	amt := 50

	if now.Day() == 6 {
		amt = 250
	} else if now.Day() == 24 {
		amt = 1000
	}

	codes = append(codes, code.GenerateCode(amt, 24, 0))

	fields := []*discordgo.MessageEmbedField{}

	for i, code := range codes {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("🎁・Day %d", i+1),
			Value:  fmt.Sprintf("||%s||", code),
			Inline: true,
		})
	}

	for i := len(codes); i < 24; i++ {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("🎁・Day %d", i+1),
			Value:  "||#######||",
			Inline: true,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:  "🎄・Advent Calendar",
		Fields: fields,
		Color:  0xB54DFF,
	}

	s.ChannelMessageEditEmbed(channel.ID, msg.ID, embed)

	if now.Day() == 24 {
		setup = false
		codes = []string{}
	}
}
