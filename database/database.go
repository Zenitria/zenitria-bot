package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CTX       = context.TODO()
	DiscordDB *mongo.Database
	GetXNODB  *mongo.Database
	GetBANDB  *mongo.Database
	Client    *mongo.Client
	TxnOpts   = options.Transaction().SetWriteConcern(writeconcern.Majority()).SetReadConcern(readconcern.Majority())
)

func Connect(uri string, db string) *mongo.Database {
	opts := options.Client().ApplyURI(uri)

	var err error
	Client, err = mongo.Connect(CTX, opts)

	if err != nil {
		panic(err)
	}

	err = Client.Ping(CTX, nil)

	if err != nil {
		panic("Connection error: " + err.Error())
	}

	return Client.Database(db)
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
		Cash:        0,
	}
}

func NewCode(code string, amt int, exp time.Time, uses int) *Code {
	return &Code{
		Code:      code,
		Amount:    amt,
		ExpiresAt: exp,
		Uses:      uses,
		Used:      0,
		Users:     []string{},
		IPs:       []string{},
	}
}
