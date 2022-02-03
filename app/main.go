package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	err := godotenv.Load(".env")
	if err != nil {
		// log.Println("unread .env")
	}

	portHTTP := os.Getenv("PORT")
	portHTTPS := os.Getenv("PORTHTTPS")

	// runServer(portHTTP, portHTTPS)
}
