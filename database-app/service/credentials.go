package service

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Lshortfile)
	if err := godotenv.Load("./../.env"); err == nil {
		PSQLHost = os.Getenv("POSTGRES_HOST")
		PSQLPort = os.Getenv("POSTGRES_PORT")
		PSQLUser = os.Getenv("POSTGRES_USERNAME")
		PSQLPassword = os.Getenv("POSTGRES_PASSWORD")
		PSQLDBName = os.Getenv("POSTGRES_DB")
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", PSQLHost, PSQLPort, PSQLUser, PSQLPassword, PSQLDBName, PSQLSSL)
	}
}

var (
	PSQLHost     = os.Getenv("POSTGRES_HOST")
	PSQLPort     = os.Getenv("POSTGRES_PORT")
	PSQLUser     = os.Getenv("POSTGRES_USERNAME")
	PSQLPassword = os.Getenv("POSTGRES_PASSWORD")
	PSQLDBName   = os.Getenv("POSTGRES_DB")
	PSQLSSL      = "disable"
)

const dbDriver = "postgres"

var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", PSQLHost, PSQLPort, PSQLUser, PSQLPassword, PSQLDBName, PSQLSSL)
