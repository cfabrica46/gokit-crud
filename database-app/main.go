package main

import (
	"log"
	"os"

	"github.com/cfabrica46/gokit-crud/database-app/service"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env loaded")
	}

	service.Run(os.Getenv("PORT"))
}
