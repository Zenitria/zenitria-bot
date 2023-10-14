package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	TOKEN           string
	MONGODB_URI     string
	GET_XNO_API_URL string
)

func init() {
	godotenv.Load()

	TOKEN = os.Getenv("TOKEN")
	MONGODB_URI = os.Getenv("MONGODB_URI")
	GET_XNO_API_URL = os.Getenv("GET_XNO_API_URL")
}
