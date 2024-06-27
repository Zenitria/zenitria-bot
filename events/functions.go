package events

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
	"zenitria-bot/prices"

	"github.com/bwmarrin/discordgo"
)

var (
	weeklyCodeScheduler *gocron.Scheduler
)

func updateStatus(s *discordgo.Session) {
	for {
		statuses := []string{
			fmt.Sprintf("BTC: $%.2f (%s)", prices.Prices.BTC.Price, change(prices.Prices.BTC.Change)),
			fmt.Sprintf("BAN: $%.2f (%s)", prices.Prices.BAN.Price, change(prices.Prices.BAN.Change)),
			fmt.Sprintf("XMR: $%.2f (%s)", prices.Prices.XMR.Price, change(prices.Prices.XMR.Change)),
			fmt.Sprintf("XNO: $%.2f (%s)", prices.Prices.XNO.Price, change(prices.Prices.XNO.Change)),
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

func removeWeeklyCodeCron() {
	if weeklyCodeScheduler != nil {
		weeklyCodeScheduler.Clear()
		weeklyCodeScheduler.Stop()
		weeklyCodeScheduler = nil
	}
}
