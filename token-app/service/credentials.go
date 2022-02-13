package service

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	var err error
	log.SetFlags(log.Lshortfile)
	if err = godotenv.Load(".env"); err != nil {
		log.Println(".env not loaded")
	}

	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")

	/* KeyData, err = ioutil.ReadFile(KeyFile)
	if err != nil {
		log.Fatal(err)
	} */
}

var (
	RedisHost string
	RedisPort string

	KeyFile = "server.key"
	// KeyData []byte
)
