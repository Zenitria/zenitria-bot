package codes

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"time"
	"zenitria-bot/database"
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

func createExpires(hrs int) time.Time {
	now := time.Now().UTC()
	expires := now.Add(time.Hour * time.Duration(hrs))

	return expires
}

func checkCode(code string) bool {
	getXNOColl := database.GetXNODB.Collection("Codes")
	getBANColl := database.GetBANDB.Collection("Codes")

	getXNOErr := getXNOColl.FindOne(database.CTX, bson.M{"codes": code}).Err()
	getBANErr := getBANColl.FindOne(database.CTX, bson.M{"codes": code}).Err()

	return errors.Is(getXNOErr, mongo.ErrNoDocuments) && errors.Is(getBANErr, mongo.ErrNoDocuments)
}

func addCode(code string, amt int, exp time.Time, uses int) {
	getXNOColl := database.GetXNODB.Collection("Codes")
	getBANColl := database.GetBANDB.Collection("Codes")

	newCode := database.NewCode(code, amt, exp, uses)

	getXNOColl.InsertOne(database.CTX, newCode)
	getBANColl.InsertOne(database.CTX, newCode)
}
