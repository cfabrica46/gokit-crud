package service

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Lshortfile)
	err := godotenv.Load(".env")
	if err == nil {
		log.Println(".env loaded")
	}

	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
}

var (
	RedisHost string
	RedisPort string

	KeyFile = "server.key"
	// KeyData []byte
)
