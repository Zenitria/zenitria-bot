package commands

import (
	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(s *discordgo.Session) {
	slashCommands := []*discordgo.ApplicationCommand{
		// General
		{
			Name:        "user-info",
			Description: "Get information about yourself or another user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to get information about",
					Required:    false,
				},
			},
		},
		{
			Name:        "server-info",
			Description: "Get information about the server",
		},
		{
			Name:        "get-xno",
			Description: "Get the stats of Get XNO",
		},
		{
			Name:        "help",
			Description: "Get help with the bot",
		},
		// Leveling
		{
			Name:        "rank",
			Description: "Get your rank or the rank of another user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to get the rank of",
					Required:    false,
				},
			},
		},
		{
			Name:        "leaderboard",
			Description: "Get the server's leaderboard",
		},
		{
			Name:        "excluded-channels",
			Description: "List all excluded channels",
		},
		{
			Name:        "exclude-channel",
			Description: "Exclude a channel from the leveling system",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel",
					Description: "The channel to exclude",
					Required:    true,
				},
			},
		},
		{
			Name:        "include-channel",
			Description: "Include a channel in the leveling system",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel",
					Description: "The channel to include",
					Required:    true,
				},
			},
		},
		// Moderation
		{
			Name:        "ban",
			Description: "Bans a user from the server",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to ban",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the ban",
					Required:    false,
				},
			},
		},
		{
			Name:        "unban",
			Description: "Unbans a user from the server",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "The user ID to unban",
				},
			},
		},
		{
			Name:        "kick",
			Description: "Kicks a user from the server",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to kick",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the kick",
					Required:    false,
				},
			},
		},
		{
			Name:        "timeout",
			Description: "Timeout a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to timeout",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "duration",
					Description: "The duration of the timeout",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "60 seconds",
							Value: 60,
						},
						{
							Name:  "5 minutes",
							Value: 300,
						},
						{
							Name:  "10 minutes",
							Value: 600,
						},
						{
							Name:  "1 hour",
							Value: 3600,
						},
						{
							Name:  "1 day",
							Value: 86400,
						},
						{
							Name:  "1 week",
							Value: 604800,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the timeout",
					Required:    false,
				},
			},
		},
	}

	s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", slashCommands)
}
