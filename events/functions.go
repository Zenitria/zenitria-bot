package events

import (
	"fmt"
	"time"
	"zenitria-bot/coingecko"

	"github.com/bwmarrin/discordgo"
)

func updateStatus(s *discordgo.Session) {
	for {
		statuses := []string{
			fmt.Sprintf("BTC: $%.2f (%s)", coingecko.Prices.Bitcoin.Price, change(coingecko.Prices.Bitcoin.Change)),
			fmt.Sprintf("XMR: $%.2f (%s)", coingecko.Prices.Monero.Price, change(coingecko.Prices.Monero.Change)),
			fmt.Sprintf("XNO: $%.2f (%s)", coingecko.Prices.Nano.Price, change(coingecko.Prices.Nano.Change)),
		}

		for _, st := range statuses {
			s.UpdateStatusComplex(
				discordgo.UpdateStatusData{
					Activities: []*discordgo.Activity{
						{
							Name: st,
							Type: discordgo.ActivityTypeWatching,
						},
					},
					Status: string(discordgo.StatusDoNotDisturb),
				},
			)

			time.Sleep(5 * time.Second)
		}
	}
}

func change(c float64) string {
	if c > 0 {
		return fmt.Sprintf("+%.2f", c) + "%"
	}

	return fmt.Sprintf("%.2f", c) + "%"
}
