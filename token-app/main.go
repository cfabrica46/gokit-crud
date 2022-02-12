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

	runServer(os.Getenv("PORT"))
}

func runServer(port string) {
}
