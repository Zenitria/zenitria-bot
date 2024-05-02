package economy

import (
	"zenitria-bot/manager"
)

func checkBalance(id string, am float32) bool {
	user := manager.GetUser(id)

	return user.Cash >= am
}

func updateBalance(id string, am float32) {
	user := manager.GetUser(id)

	user.Cash += am

	manager.UpdateUser(id, user.Level, user.XP, user.NextLevelXP, user.Warnings, user.Cash, user.LastClaimed)
}
