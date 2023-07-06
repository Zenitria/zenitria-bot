package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CTX       = context.TODO()
	DiscordDB *mongo.Database
	GetXNODB  *mongo.Database
)

func Connect(uri string, db string) *mongo.Database {
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(CTX, opts)

	if err != nil {
		panic(err)
	}

	err = client.Ping(CTX, nil)

	if err != nil {
		panic("Connection error: " + err.Error())
	}

	return client.Database(db)
}

func Disconnect(db *mongo.Database) error {
	err := db.Client().Disconnect(CTX)

	return err
}

func NewUser(id string) *User {
	return &User{
		ID:          id,
		Level:       0,
		XP:          0,
		NextLevelXP: 100,
		Warnings:    0,
	}
}

func NewCode(code string, amt int, exp string, uses int) *Code {
	return &Code{
		Code:    code,
		Amount:  amt,
		Expires: exp,
		Uses:    uses,
		Used:    0,
		Users:   []string{},
		IPs:     []string{},
	}
}
