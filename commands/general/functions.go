package general

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func getSlashCommandMentions(s *discordgo.Session) map[string]string {
	cmds, _ := s.ApplicationCommands(s.State.User.ID, "")

	mentions := make(map[string]string)

	for _, cmd := range cmds {
		mentions[cmd.Name] = fmt.Sprintf("</%s:%s>", cmd.Name, cmd.ID)
	}

	return mentions
}

func getSlashSubcommandMentions(s *discordgo.Session) map[string]string {
	cmds, _ := s.ApplicationCommands(s.State.User.ID, "")

	mentions := make(map[string]string)

	for _, cmd := range cmds {
		for _, option := range cmd.Options {
			if option.Type == discordgo.ApplicationCommandOptionSubCommand {
				mentions[fmt.Sprintf("%s %s", cmd.Name, option.Name)] = fmt.Sprintf("</%s %s:%s>", cmd.Name, option.Name, cmd.ID)
			}
		}
	}

	return mentions
}
