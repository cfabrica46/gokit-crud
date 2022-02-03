package service

import "fmt"

const (
	PSQLHost     = "localhost"
	PSQLPort     = 5431
	PSQLUser     = "cfabrica46"
	PSQLPassword = "01234"
	PSQLDBName   = "go_crud"
	PSQLSSL      = "require"
)

const dbDriver = "postgres"

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", PSQLHost, PSQLPort, PSQLUser, PSQLPassword, PSQLDBName, PSQLSSL)
