package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	TOKEN               string
	MONGODB_URI         string
	GET_XNO_API_URL     string
	COMMANDS_CHANNEL_ID string
	OWNER_ID            string
)

func init() {
	godotenv.Load()

	TOKEN = os.Getenv("TOKEN")
	MONGODB_URI = os.Getenv("MONGODB_URI")
	GET_XNO_API_URL = os.Getenv("GET_XNO_API_URL")
	COMMANDS_CHANNEL_ID = os.Getenv("COMMANDS_CHANNEL_ID")
	OWNER_ID = os.Getenv("OWNER_ID")
}
