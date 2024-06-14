package events

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
	"zenitria-bot/codes"
	"zenitria-bot/config"
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

func sendWeeklyCode(s *discordgo.Session) {
	code, expires := codes.GenerateCode(50, 24, 0)

	embed := &discordgo.MessageEmbed{
		Title:       "💎・Diamonds Code",
		Description: fmt.Sprintf("🏷️・**Code:** %s\n💎・**Diamonds:** 50\n⏳・**Expires:** <t:%d:R>\n💰・**Redeem:** [Get XNO](https://get-xno.com/app/redeem) & [Get BAN](https://get-ban.com/app/redeem)", code, expires.Unix()),
		Color:       0xB54DFF,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/SAJ5PrWD0DcAAAAC/diamond.gif",
		},
	}

	s.ChannelMessageSendEmbed(config.CODES_CHANNEL_ID, embed)
	s.ChannelMessageSend(config.CODES_CHANNEL_ID, fmt.Sprintf("<@&%s>", config.CODES_ROLE_ID))
}

func weeklyCode(s *discordgo.Session) {
	weeklyCodeScheduler = gocron.NewScheduler(time.UTC)
	weeklyCodeScheduler.Cron("0 18 * * 5").Do(sendWeeklyCode, s)
	weeklyCodeScheduler.StartAsync()
}

func removeWeeklyCodeCron() {
	if weeklyCodeScheduler != nil {
		weeklyCodeScheduler.Clear()
		weeklyCodeScheduler.Stop()
		weeklyCodeScheduler = nil
	}
}
