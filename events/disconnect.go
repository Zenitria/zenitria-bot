package events

import (
	"github.com/bwmarrin/discordgo"
)

func Disconnenct(s *discordgo.Session, e *discordgo.Disconnect) {
	removeWeeklyCodeCron()
}
