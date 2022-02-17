package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env loaded")
	}

	runServer(os.Getenv("PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("TOKEN_HOST"), os.Getenv("TOKEN_PORT"))
}

func runServer(port, dbHost, dbPort, tokenHost, tokenPort string) {
}
