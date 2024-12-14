package events

import (
	"fmt"
	"strings"
	"time"
	"zenitria-bot/codes"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)

var (
	generatedCodes []string
	setup          = false
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
		Color:  0xBE4DFF,
	}

	msg, _ := s.ChannelMessageSendEmbed(channel.ID, embed)

	embed = &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("The advent calendar has been set on %s channel.", channel.Mention()),
		Color:       0xBE4DFF,
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

func HandleFixAdventCalendar(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckOwner(s, i) {
		return
	}

	data := i.ApplicationCommandData()

	var channel *discordgo.Channel
	var msgId string
	var oldCodes string

	for _, option := range data.Options {
		switch option.Name {
		case "channel":
			channel = option.ChannelValue(s)
		case "message_id":
			msgId = option.StringValue()
		case "codes":
			oldCodes = option.StringValue()
		}
	}

	split := strings.Split(oldCodes, ",")

	for _, code := range split {
		generatedCodes = append(generatedCodes, strings.Replace(code, " ", "", -1))
	}

	now := time.Now().UTC()

	amt := 50

	if now.Day() == 6 {
		amt = 250
	} else if now.Day() == 24 {
		amt = 1000
	}

	code, _ := codes.GenerateCode(amt, 24, 0)

	generatedCodes = append(generatedCodes, code)

	var fields []*discordgo.MessageEmbedField

	for i, code := range generatedCodes {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("🎁・Day %d", i+1),
			Value:  fmt.Sprintf("||%s||", code),
			Inline: true,
		})
	}

	for i := len(generatedCodes); i < 24; i++ {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("🎁・Day %d", i+1),
			Value:  "||#######||",
			Inline: true,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:  "🎄・Advent Calendar",
		Fields: fields,
		Color:  0xBE4DFF,
	}

	msg, _ := s.ChannelMessageEditEmbed(channel.ID, msgId, embed)

	if now.Day() == 24 {
		setup = false
		generatedCodes = []string{}
	}

	embed = &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("The advent calendar has been fixed on %s channel.", channel.Mention()),
		Color:       0xBE4DFF,
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

	scheduler.Cron("0 0 * * *").LimitRunsTo(24 - now.Day()).Do(func() {
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

	code, _ := codes.GenerateCode(amt, 24, 0)

	generatedCodes = append(generatedCodes, code)

	var fields []*discordgo.MessageEmbedField

	for i, code := range generatedCodes {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("🎁・Day %d", i+1),
			Value:  fmt.Sprintf("||%s||", code),
			Inline: true,
		})
	}

	for i := len(generatedCodes); i < 24; i++ {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("🎁・Day %d", i+1),
			Value:  "||#######||",
			Inline: true,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:  "🎄・Advent Calendar",
		Fields: fields,
		Color:  0xBE4DFF,
	}

	s.ChannelMessageEditEmbed(channel.ID, msg.ID, embed)

	if now.Day() == 24 {
		setup = false
		generatedCodes = []string{}
	}
}
