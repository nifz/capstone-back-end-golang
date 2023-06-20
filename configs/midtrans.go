package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMidtransServerKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MIDTRANS_SERVER_KEY")
}

func EnvMidtransClientKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MIDTRANS_CLIENT_KEY")
}
