package economy

import (
	"zenitria-bot/usermanager"
)

func checkBalance(id string, price float32) bool {
	user := usermanager.GetUser(id)

	return user.Cash >= price
}

func updateBalance(id string, price float32) {
	user := usermanager.GetUser(id)

	user.Cash += price

	usermanager.UpdateUser(id, user.Level, user.XP, user.NextLevelXP, user.Warnings, user.Cash)
}
