package main

import (
	"os"

	"github.com/cfabrica46/gokit-crud/database-app/service"
)

func main() {
	service.RunServer(os.Getenv("PORT"))
}
