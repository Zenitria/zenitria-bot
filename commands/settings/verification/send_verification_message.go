package verification

import (
	"fmt"
	"zenitria-bot/manager"

	"github.com/bwmarrin/discordgo"
)

func HandleSend(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	channel := data.Options[0].Options[0].ChannelValue(s)
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
