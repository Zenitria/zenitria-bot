package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func pointer[t any](v t) *t {
	return &v
}

func RegisterCommands(s *discordgo.Session) {
	slashCommands := []*discordgo.ApplicationCommand{
		// General
		// General/Info
		{
			Name:         "info",
			Description:  "Get information about something",
			DMPermission: pointer(false),
			Options: []*discordgo.ApplicationCommandOption{
				// General/Info/User
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "user",
					Description: "Get information about a user",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionUser,
							Name:        "user",
							Description: "The user to get information about",
							Required:    false,
						},
					},
				},
				// General/Info/Server
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "server",
					Description: "Get information about the server",
				},
			},
		},
		// General/Stats
		{
			Name:         "stats",
			Description:  "Get the stats of something",
			DMPermission: pointer(false),
			Options: []*discordgo.ApplicationCommandOption{
				// General/Stats/Get-XNO
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "get-xno",
					Description: "Get the stats of Get XNO",
				},
				// General/Stats/Get-BAN
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "get-ban",
					Description: "Get the stats of Get BAN",
				},
			},
		},
		// General/Help
		{
			Name:         "help",
			Description:  "Get help with the bot",
			DMPermission: pointer(false),
		},
		// Leveling
		// Leveling/Rank
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
		// Leveling/Leaderboard
		{
			Name:         "leaderboard",
			Description:  "Get the server's leaderboard",
			DMPermission: pointer(false),
		},
		// Earning
		// Earning/Balance
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
		// Earning/Withdraw
		{
			Name:         "withdraw",
			Description:  "Withdraw your balance",
			DMPermission: pointer(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "amount",
					Description: "The amount to withdraw",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "cryptocurrency",
					Description: "The cryptocurrency to withdraw",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Nano (XNO)",
							Value: "XNO",
						},
						{
							Name:  "Banano (BAN)",
							Value: "BAN",
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "wallet_address",
					Description: "The wallet address to withdraw to",
					Required:    true,
				},
			},
		},
		// Earning/Claim
		{
			Name:         "claim",
			Description:  "Claim your 15 minute reward",
			DMPermission: pointer(false),
		},
		// Earning/Balances
		{
			Name:         "balances",
			Description:  "Get the bot crypto balances.",
			DMPermission: pointer(false),
		},
		// Earning/Rich List
		{
			Name:         "rich-list",
			Description:  "Get the server's rich list",
			DMPermission: pointer(false),
		},
		// Earning/Add Dollars
		{
			Name:                     "add-dollars",
			Description:              "Add dollars to a user",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionAdministrator),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to add dollars to",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "amount",
					Description: "The amount of dollars to add",
					Required:    true,
				},
			},
		},
		// Moderation
		// Moderation/Ban
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
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "delete_messages",
					Description: "The part of the user's recent message history that you want to delete",
					Required:    false,
					MinValue:    pointer[float64](1),
					MaxValue:    7,
				},
			},
		},
		// Moderation/Unban
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
		// Moderation/Kick
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
		// Moderation/Timeout
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
		// Moderation/Warn
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
		// Moderation/Warns
		{
			Name:         "warns",
			Description:  "Get your warns or the warns of another user",
			DMPermission: pointer(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to get the warns of",
				},
			},
		},
		// Settings
		// Settings/Leveling
		{
			Name:                     "leveling",
			Description:              "Manage the leveling system",
			DMPermission:             pointer(false),
			DefaultMemberPermissions: pointer[int64](discordgo.PermissionManageChannels),
			Options: []*discordgo.ApplicationCommandOption{
				// Settings/Leveling/Exclude
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "exclude",
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
				// Settings/Leveling/Include
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "include",
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
				// Settings/Leveling/Excluded-List
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "excluded-list",
					Description: "Get the list of excluded channels",
				},
			},
		},
	}

	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", slashCommands)

	if err != nil {
		fmt.Println(err)
	}
}
