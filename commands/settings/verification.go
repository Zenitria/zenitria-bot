package settings

import (
	"fmt"
	"math/rand"
	"time"
	"zenitria-bot/commands/settings/verification"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleVerification(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	handlers := map[string](func(*discordgo.Session, *discordgo.InteractionCreate)){
		"role": verification.HandleRole,
		"send": verification.HandleSend,
	}

	if handler, ok := handlers[data.Options[0].Name]; ok {
		handler(s, i)
	}
}

func HandleVerifyButton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	role, _ := manager.GetVerificationRole()

	for _, r := range i.Member.Roles {
		if r == role {
			embed := &discordgo.MessageEmbed{
				Title:       "🚫・Error!",
				Description: "You are already verified!",
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
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var first, second int
	var operator string

	operatorCode := rng.Intn(2)

	if operatorCode == 0 {
		first = rng.Intn(150) + 1
		second = rng.Intn(first) + 1
		operator = "-"
	} else {
		first = rng.Intn(150) + 1
		second = rng.Intn(150) + 1
		operator = "+"
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "verification-modal",
			Title:    "🔒・Verification!",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    fmt.Sprintf("%d|%d|%d", first, second, operatorCode),
							Label:       "Captcha",
							Style:       discordgo.TextInputShort,
							Placeholder: fmt.Sprintf("%d %s %d = ?", first, operator, second),
							Required:    true,
						},
					},
				},
			},
		},
	}

	s.InteractionRespond(i.Interaction, response)
}
