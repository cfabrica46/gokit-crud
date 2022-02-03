package service

import (
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
)

func RunServer() {
	svc := &serviceDB{}

	err := svc.OpenDB(dbDriver, psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	getAllUsersHandler := httptransport.NewServer(
		makeGetAllUsersEndpoint(svc),
		decodeGetAllUsersRequest,
		encodeResponse,
	)

	getUserByIDHandler := httptransport.NewServer(
		makeGetUserByIDEndpoint(svc),
		decodeGetUserByIDRequest,
		encodeResponse,
	)

	getUserByUsernameAndPasswordHandler := httptransport.NewServer(
		makeGetUserByUsernameAndPasswordEndpoint(svc),
		decodeGetUserByUsernameAndPasswordRequest,
		encodeResponse,
	)

	getIDByUsernameHandler := httptransport.NewServer(
		makeGetIDByUsernameEndpoint(svc),
		decodeGetIDByUsernameRequest,
		encodeResponse,
	)

	insertUserHandler := httptransport.NewServer(
		makeInsertUserEndpoint(svc),
		decodeInsertUserRequest,
		encodeResponse,
	)

	deleteUserByUsernameHandler := httptransport.NewServer(
		makeDeleteUserByUsernameEndpoint(svc),
		decodeDeleteUserByUsernameRequest,
		encodeResponse,
	)

	http.Handle("/users", getAllUsersHandler)
	http.Handle("/userByID", getUserByIDHandler)
	http.Handle("/userByUsernameAndPassword", getUserByUsernameAndPasswordHandler)
	http.Handle("/idByUsername", getIDByUsernameHandler)
	http.Handle("/insert", insertUserHandler)
	http.Handle("/delete", deleteUserByUsernameHandler)

	log.Println("ListenAndServe on localhost:8080")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
