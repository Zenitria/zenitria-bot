package moderation

import "zenitria-bot/usermanager"

func getDurationString(d int64) string {
	var output string

	switch d {
	case 60:
		output = "60 seconds"
	case 300:
		output = "5 minutes"
	case 600:
		output = "10 minutes"
	case 3600:
		output = "1 hour"
	case 86400:
		output = "1 day"
	case 604800:
		output = "1 week"
	}

	return output
}

func getWarns(id string) int {
	user := usermanager.GetUser(id)

	return user.Warnings
}

func addWarn(id string) {
	user := usermanager.GetUser(id)

	user.Warnings++

	usermanager.UpdateUser(id, user.Level, user.XP, user.NextLevelXP, user.Warnings, user.Cash)
}
