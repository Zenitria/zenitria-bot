package code

import (
	"math/rand"
	"time"
	"zenitria-bot/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func createCode() string {
	chars := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	code := make([]rune, 7)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range code {
		code[i] = chars[rng.Intn(len(chars))]
	}

	return string(code)
}

func createExpires(hrs int) string {
	now := time.Now().UTC()
	expires := now.Add(time.Hour * time.Duration(hrs)).Format("2006-01-02T15:04:05Z")

	return expires
}

func checkCode(code string) bool {
	collection := database.GetXNODB.Collection("Codes")

	err := collection.FindOne(database.CTX, bson.M{"code": code}).Err()

	return err == mongo.ErrNoDocuments
}

func addCode(code string, amt int, exp string, uses int) {
	collection := database.GetXNODB.Collection("Codes")

	newCode := database.NewCode(code, amt, exp, uses)

	collection.InsertOne(database.CTX, newCode)
}
