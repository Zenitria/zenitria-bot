package database

import "time"

type User struct {
	ID          string    `bson:"_id"`
	Level       int       `bson:"level"`
	XP          int       `bson:"xp"`
	NextLevelXP int       `bson:"nextLevelXP"`
	Warnings    int       `bson:"warnings"`
	Cash        float32   `bson:"cash"`
	LastClaimed time.Time `bson:"lastClaimed"`
}

type Code struct {
	Code      string    `bson:"code"`
	Amount    int       `bson:"amount"`
	ExpiresAt time.Time `bson:"expiresAt"`
	Uses      int       `bson:"uses"`
	Used      int       `bson:"used"`
	Users     []string  `bson:"users"`
	IPs       []string  `bson:"ips"`
}

type Channel struct {
	ID string `bson:"_id"`
}

type Setting struct {
	Name  string `bson:"_id"`
	Value any    `bson:"value"`
}
