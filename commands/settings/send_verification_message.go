package settings

import (
	"fmt"
	"math/rand"
	"time"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleSendVerificationMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	channel := data.Options[0].ChannelValue(s)
	_, err := manager.GetVerificationRole()

	if err != nil {
		cmds, _ := s.ApplicationCommands(s.State.User.ID, "")

		mention := ""

		for _, cmd := range cmds {
			if cmd.Name == "set-verification-role" {
				mention = fmt.Sprintf("</%s:%s>", cmd.Name, cmd.ID)
				break
			}
		}

		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: fmt.Sprintf("The verification role has not been set! Please set it using %s.", mention),
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
		Title:       "🔒・Verification!",
		Description: "Click the button below to verify yourself.",
		Color:       0x06e386,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/9PoGus-5E64AAAAC/budder-get-verified.gif",
		},
	}

	components := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label:    "Verify",
					Style:    discordgo.PrimaryButton,
					CustomID: "verify-button",
				},
			},
		},
	}

	message := &discordgo.MessageSend{
		Embed:      embed,
		Components: components,
	}

	_, err = s.ChannelMessageSendComplex(channel.ID, message)

	if err != nil {
		embed = &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "Enter a valid channel.",
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

	embed = &discordgo.MessageEmbed{
		Title:       "✅・Success!",
		Description: fmt.Sprintf("The verification message has been sent on %s channel.", channel.Mention()),
		Color:       0x06e386,
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
