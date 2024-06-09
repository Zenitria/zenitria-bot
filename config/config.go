package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	TOKEN               string
	MONGODB_URI         string
	GET_XNO_API_URL     string
	GET_BAN_API_URL     string
	ZENITRIA_API_URL    string
	COMMANDS_CHANNEL_ID string
	CODES_CHANNEL_ID    string
	CODES_ROLE_ID       string
	OWNER_ID            string
	NANSWAP_SECRET      string
	NANO_SEED           string
)

func init() {
	godotenv.Load()

	TOKEN = os.Getenv("TOKEN")
	MONGODB_URI = os.Getenv("MONGODB_URI")
	GET_XNO_API_URL = os.Getenv("GET_XNO_API_URL")
	GET_BAN_API_URL = os.Getenv("GET_BAN_API_URL")
	ZENITRIA_API_URL = os.Getenv("ZENITRIA_API_URL")
	COMMANDS_CHANNEL_ID = os.Getenv("COMMANDS_CHANNEL_ID")
	CODES_CHANNEL_ID = os.Getenv("CODES_CHANNEL_ID")
	CODES_ROLE_ID = os.Getenv("CODES_ROLE_ID")
	OWNER_ID = os.Getenv("OWNER_ID")
	NANSWAP_SECRET = os.Getenv("NANSWAP_SECRET")
	NANO_SEED = os.Getenv("NANO_SEED")
}
