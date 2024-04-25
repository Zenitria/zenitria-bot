package manager

import (
	"time"
	"zenitria-bot/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(id string) {
	collection := database.DiscordDB.Collection("Users")
	user := database.NewUser(id)

	collection.InsertOne(database.CTX, user)
}

func GetUser(id string) database.User {
	collection := database.DiscordDB.Collection("Users")

	var user database.User

	collection.FindOne(database.CTX, bson.M{"_id": id}).Decode(&user)

	return user
}

func CheckUser(id string) bool {
	collection := database.DiscordDB.Collection("Users")

	var user database.User

	err := collection.FindOne(database.CTX, bson.M{"_id": id}).Decode(&user)

	return err != mongo.ErrNoDocuments
}

func UpdateUser(id string, l int, xp int, n int, w int, c float32, lc time.Time) {
	session, _ := database.Client.StartSession()

	defer session.EndSession(database.CTX)

	session.WithTransaction(database.CTX, func(ctx mongo.SessionContext) (any, error) {
		collection := database.DiscordDB.Collection("Users")

		update := bson.M{
			"$set": database.User{
				ID:          id,
				Level:       l,
				XP:          xp,
				NextLevelXP: n,
				Warnings:    w,
				Cash:        c,
				LastClaimed: lc,
			},
		}

		collection.UpdateByID(database.CTX, id, update)

		return nil, nil
	})
}
