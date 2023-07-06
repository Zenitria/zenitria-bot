package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, e *discordgo.Ready) {
	go updateStatus(s)

	fmt.Println(s.State.User.String() + " is ready!")
}
