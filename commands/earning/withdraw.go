package earning

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zenitria/bananogo"
	"github.com/zenitria/nanogo"
	"math/big"
	"zenitria-bot/config"
	"zenitria-bot/manager"
	"zenitria-bot/prices"
)

func HandleWithdraw(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if manager.CheckCommandChannel(s, i, i.ChannelID) {
		return
	}

	data := i.ApplicationCommandData()

	if len(data.Options) < 3 {
		return
	}

	amount := data.Options[0].FloatValue()
	crypto := data.Options[1].StringValue()
	wallet := data.Options[2].StringValue()

	if !checkBalance(i.Member.User.ID, amount) {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "You don't have enough money to withdraw.",
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

	if amount < 0.01 {
		embed := &discordgo.MessageEmbed{
			Title:       "🚫・Error!",
			Description: "The minimum amount to withdraw is $0.01.",
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

	if crypto == "XNO" {
		if !nanogo.AddressIsValid(wallet) {
			embed := &discordgo.MessageEmbed{
				Title:       "🚫・Error!",
				Description: "The wallet address is invalid.",
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
			Title:       "💸・Withdraw",
			Description: "Withdraw in progress...",
			Color:       0xBE4DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://i.gifer.com/90JG.gif",
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		}

		s.InteractionRespond(i.Interaction, response)

		updateBalance(i.Member.User.ID, -amount)

		c := nanogo.Client{
			Url:        "https://nodes.nanswap.com/XNO",
			AuthHeader: "nodes-api-key",
			AuthToken:  config.NANSWAP_SECRET,
		}

		raw, err := nanogo.NanoToRaw(fmt.Sprintf("%f", amount/prices.Prices.XNO.Price))

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		privkey, err := nanogo.SeedToPrivateKey(config.NANO_SEED, 0)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		pubkey, err := nanogo.PrivateKeyToPublicKey(privkey)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		addr, err := nanogo.PublicKeyToAddress(pubkey)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		bals, err := c.GetAccountBalance(addr)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		bal := new(big.Int)

		if _, ok := bal.SetString(bals.Balance, 10); !ok {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		rawBigInt := new(big.Int)

		if _, ok := rawBigInt.SetString(raw, 10); !ok {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		if bal.Cmp(rawBigInt) == -1 {
			updateBalance(i.Member.User.ID, amount)

			embed := &discordgo.MessageEmbed{
				Title:       "🚫・Error!",
				Description: "You don't have enough balance to withdraw.",
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
			return
		}

		hash, err := c.Send(wallet, raw, config.NANO_SEED, 0)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		sendSuccess(s, i, amount, hash, "https://blocklattice.io/block/")
		return
	} else if crypto == "BAN" {
		if !bananogo.AddressIsValid(wallet) {
			embed := &discordgo.MessageEmbed{
				Title:       "🚫・Error!",
				Description: "The wallet address is invalid.",
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
			Title:       "💸・Withdraw",
			Description: "Withdraw in progress...",
			Color:       0xBE4DFF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://i.gifer.com/90JG.gif",
			},
		}

		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		}

		s.InteractionRespond(i.Interaction, response)

		updateBalance(i.Member.User.ID, -amount)

		c := bananogo.Client{
			Url:        "https://nodes.nanswap.com/BAN",
			AuthHeader: "nodes-api-key",
			AuthToken:  config.NANSWAP_SECRET,
		}

		raw, err := bananogo.BananoToRaw(fmt.Sprintf("%f", amount/prices.Prices.BAN.Price))

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		privkey, err := bananogo.SeedToPrivateKey(config.NANO_SEED, 0)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		pubkey, err := bananogo.PrivateKeyToPublicKey(privkey)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		addr, err := bananogo.PublicKeyToAddress(pubkey)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		bals, err := c.GetAccountBalance(addr)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		bal := new(big.Int)

		if _, ok := bal.SetString(bals.Balance, 10); !ok {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		rawBigInt := new(big.Int)

		if _, ok := rawBigInt.SetString(raw, 10); !ok {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		if bal.Cmp(rawBigInt) == -1 {
			updateBalance(i.Member.User.ID, amount)

			embed := &discordgo.MessageEmbed{
				Title:       "🚫・Error!",
				Description: "You don't have enough balance to withdraw.",
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
			return
		}

		hash, err := c.Send(wallet, raw, config.NANO_SEED, 0)

		if err != nil {
			updateBalance(i.Member.User.ID, amount)
			sendInternalError(s, i)
			return
		}

		sendSuccess(s, i, amount, hash, "https://creeper.banano.cc/hash/")
		return
	}
}
