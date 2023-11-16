package manager

import (
	"fmt"
	"zenitria-bot/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsChannelExcluded(ch string) bool {
	collection := database.DiscordDB.Collection("Settings")

	var setting database.Setting

	err := collection.FindOne(database.CTX, bson.M{"_id": "Excluded Channels"}).Decode(&setting)

	if err != nil {
		return false
	}

	for _, channel := range setting.Value.(primitive.A) {
		if channel == ch {
			return true
		}
	}

	return false
}

func ExcludeChannel(ch string) {
	collection := database.DiscordDB.Collection("Settings")

	err := collection.FindOne(database.CTX, bson.M{"_id": "Excluded Channels"}).Err()

	if err != nil {
		collection.InsertOne(database.CTX, bson.M{"_id": "Excluded Channels", "value": []string{ch}})
		return
	}

	collection.UpdateOne(database.CTX, bson.M{"_id": "Excluded Channels"}, bson.M{"$push": bson.M{"value": ch}})
}

func IncludeChannel(ch string) {
	collection := database.DiscordDB.Collection("Settings")

	collection.UpdateOne(database.CTX, bson.M{"_id": "Excluded Channels"}, bson.M{"$pull": bson.M{"value": ch}})
}

func GetExcludedChannels() string {
	collection := database.DiscordDB.Collection("Settings")

	var setting database.Setting

	err := collection.FindOne(database.CTX, bson.M{"_id": "Excluded Channels"}).Decode(&setting)

	if err != nil {
		return "No channels are excluded from leveling system."
	}

	var output string

	for i, ch := range setting.Value.(primitive.A) {
		output += fmt.Sprintf("%d. <#%s>\n", i+1, ch)
	}

	if output == "" {
		output = "No channels are excluded from leveling system."
	}

	return output
}
