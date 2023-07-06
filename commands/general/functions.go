package general

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func getSlashCommandMentions(s *discordgo.Session) map[string]string {
	cmds, _ := s.ApplicationCommands(s.State.User.ID, "")

	ids := make(map[string]string)

	for _, cmd := range cmds {
		ids[cmd.Name] = cmd.ID
	}

	mentions := make(map[string]string)

	for name, id := range ids {
		mentions[name] = fmt.Sprintf("</%s:%s>", name, id)
	}

	return mentions
}
