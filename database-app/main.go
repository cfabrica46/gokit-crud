package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cfabrica46/gokit-crud/database-app/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	log.SetFlags(log.Lshortfile)
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
	}

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))

	fmt.Println(dbInfo)

	db, err := sql.Open(os.Getenv("DB_DRIVER"), dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	runServer(os.Getenv("PORT"), db)
}

func runServer(port string, db *sql.DB) {
	svc := service.GetService(db)

	getAllUsersHandler := httptransport.NewServer(
		service.MakeGetAllUsersEndpoint(svc),
		service.DecodeGetAllUsersRequest,
		service.EncodeResponse,
	)

	getUserByIDHandler := httptransport.NewServer(
		service.MakeGetUserByIDEndpoint(svc),
		service.DecodeGetUserByIDRequest,
		service.EncodeResponse,
	)

	getUserByUsernameAndPasswordHandler := httptransport.NewServer(
		service.MakeGetUserByUsernameAndPasswordEndpoint(svc),
		service.DecodeGetUserByUsernameAndPasswordRequest,
		service.EncodeResponse,
	)

	getIDByUsernameHandler := httptransport.NewServer(
		service.MakeGetIDByUsernameEndpoint(svc),
		service.DecodeGetIDByUsernameRequest,
		service.EncodeResponse,
	)

	insertUserHandler := httptransport.NewServer(
		service.MakeInsertUserEndpoint(svc),
		service.DecodeInsertUserRequest,
		service.EncodeResponse,
	)

	deleteUserHandler := httptransport.NewServer(
		service.MakeDeleteUserEndpoint(svc),
		service.DecodeDeleteUserRequest,
		service.EncodeResponse,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/users").Handler(getAllUsersHandler)
	r.Methods(http.MethodGet).Path("/user/id").Handler(getUserByIDHandler)
	r.Methods(http.MethodGet).Path("/user/username_password").Handler(getUserByUsernameAndPasswordHandler)
	r.Methods(http.MethodGet).Path("/id/username").Handler(getIDByUsernameHandler)
	r.Methods(http.MethodPost).Path("/user").Handler(insertUserHandler)
	r.Methods(http.MethodDelete).Path("/user").Handler(deleteUserHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+port, r))
}
