package earning

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/zenitria/bananogo"
	"github.com/zenitria/nanogo"
	"strconv"
	"zenitria-bot/config"
	"zenitria-bot/prices"
)

func HandleBalances(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "📊・Balances",
		Description: "Please wait while we fetch balances...",
		Color:       0xBE4DFF,
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	}

	s.InteractionRespond(i.Interaction, response)

	cXno := nanogo.Client{
		Url:        "https://nodes.nanswap.com/XNO",
		AuthHeader: "nodes-api-key",
		AuthToken:  config.NANSWAP_SECRET,
	}

	cBan := bananogo.Client{
		Url:        "https://nodes.nanswap.com/BAN",
		AuthHeader: "nodes-api-key",
		AuthToken:  config.NANSWAP_SECRET,
	}

	privkey, err := nanogo.SeedToPrivateKey(config.NANO_SEED, 0)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	pubkey, err := nanogo.PrivateKeyToPublicKey(privkey)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	addr, err := nanogo.PublicKeyToAddress(pubkey)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	balsXno, err := cXno.GetAccountBalance(addr)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	balXno, err := nanogo.RawToNano(balsXno.Balance)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	balXnoFloat, err := strconv.ParseFloat(balXno, 64)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	addr, err = bananogo.PublicKeyToAddress(pubkey)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	balsBan, err := cBan.GetAccountBalance(addr)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	balBan, err := bananogo.RawToBanano(balsBan.Balance)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	balBanFloat, err := strconv.ParseFloat(balBan, 64)

	if err != nil {
		sendInternalError(s, i)
		return
	}

	embed = &discordgo.MessageEmbed{
		Title:       "📊・Balances",
		Description: "Here are our balances:",
		Color:       0xBE4DFF,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "XNO",
				Value:  fmt.Sprintf("Ӿ%s ($%s)", balXno, formatFloat(balXnoFloat*prices.Prices.XNO.Price)),
				Inline: false,
			},
			{
				Name:   "BAN",
				Value:  fmt.Sprintf("%s BAN ($%s)", balBan, formatFloat(balBanFloat*prices.Prices.BAN.Price)),
				Inline: false,
			},
		},
	}

	embeds := []*discordgo.MessageEmbed{embed}

	newResponse := &discordgo.WebhookEdit{
		Embeds: &embeds,
	}

	s.InteractionResponseEdit(i.Interaction, newResponse)
}
