package commands

import (
	"github.com/bwmarrin/discordgo"
)

func pointer[t any](v t) *t {
	return &v
}

func RegisterCommands(s *discordgo.Session) {
	slashCommands := []*discordgo.ApplicationCommand{
		// General
		{
			Name:         "user-info",
			Description:  "Get information about yourself or another user",
			DMPermission: pointer(false),
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
			Name:         "server-info",
			Description:  "Get information about the server",
			DMPermission: pointer(false),
		},
		{
			Name:         "get-xno",
			Description:  "Get the stats of Get XNO",
			DMPermission: pointer(false),
		},
		{
			Name:         "help",
			Description:  "Get help with the bot",
			DMPermission: pointer(false),
		},
		// Leveling
		{
			Name:         "rank",
			Description:  "Get your rank or the rank of another user",
			DMPermission: pointer(false),
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
			Name:         "leaderboard",
			Description:  "Get the server's leaderboard",
			DMPermission: pointer(false),
		},
		// Economy
		{
			Name:         "balance",
			Description:  "Get your balance or the balance of another user",
			DMPermission: pointer(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to get the balance of",
					Required:    false,
				},
			},
		},
		{
			Name:         "shop",
			Description:  "Buy items with your cash",
			DMPermission: pointer(false),
		},
		{
			Name:         "buy",
			Description:  "Buy an item from the shop",
			DMPermission: pointer(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "item",
					Description: "The item to buy",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "💎・Diamonds Pack: ✨・Mini (10)",
							Value: "diamonds-mini",
						},
						{
							Name:  "💎・Diamonds Pack: 🌟・Small (25)",
							Value: "diamonds-small",
						},
						{
							Name:  "💎・Diamonds Pack: 🎁・Medium (100)",
							Value: "diamonds-medium",
						},
						{
							Name:  "💎・Diamonds Pack: 🔥・Big (250)",
							Value: "diamonds-big",
						},
						{
							Name:  "💎・Diamonds Pack: 🏆・Premium (1000)",
							Value: "diamonds-premium",
						},
					},
				},
			},
		},
		{
			Name:         "claim",
			Description:  "Claim your hourly reward",
			DMPermission: pointer(false),
		},
		// Moderation
		{
			Name:                     "ban",
			Description:              "Bans a user from the server",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionBanMembers),
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
			Name:                     "unban",
			Description:              "Unbans a user from the server",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionBanMembers),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "The user ID to unban",
				},
			},
		},
		{
			Name:                     "kick",
			Description:              "Kicks a user from the server",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionKickMembers),
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
			Name:                     "timeout",
			Description:              "Timeout a user",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionModerateMembers),
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
		{
			Name:                     "warn",
			Description:              "Warn a user",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionModerateMembers),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to warn",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the warn",
					Required:    false,
				},
			},
		},
		// Settings
		{
			Name:                     "set-verification-role",
			Description:              "Set the verification role",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionAdministrator),

			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The role to set as the verification role",
					Required:    true,
				},
			},
		},
		{
			Name:                     "send-verification-message",
			Description:              "Send the verification message to selected channel",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionAdministrator),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel",
					Description: "The channel to send the verification message to",
					Required:    true,
				},
			},
		},
		{
			Name:                     "excluded-channels",
			Description:              "List all excluded channels from the leveling system",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionManageChannels),
		},
		{
			Name:                     "exclude-channel",
			Description:              "Exclude a channel from the leveling system",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionManageChannels),
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
			Name:                     "include-channel",
			Description:              "Include a channel in the leveling system",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionManageChannels),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel",
					Description: "The channel to include",
					Required:    true,
				},
			},
		},
	}

	s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", slashCommands)
}
