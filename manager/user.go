package manager

import (
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

func UpdateUser(id string, l int, xp int, n int, w int, c float32) {
	collection := database.DiscordDB.Collection("Users")

	update := bson.M{
		"$set": database.User{
			ID:          id,
			Level:       l,
			XP:          xp,
			NextLevelXP: n,
			Warnings:    w,
			Cash:        c,
		},
	}

	collection.UpdateByID(database.CTX, id, update)
}
