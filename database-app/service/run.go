package service

import (
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func Run(port string) {
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

	deleteUserHandler := httptransport.NewServer(
		makeDeleteUserEndpoint(svc),
		decodeDeleteUserRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/users").Handler(getAllUsersHandler)
	r.Methods(http.MethodGet).Path("/user/{id:[0-9]+}").Handler(getUserByIDHandler)
	r.Methods(http.MethodGet).Path("/user/username_password").Handler(getUserByUsernameAndPasswordHandler)
	r.Methods(http.MethodGet).Path("/id/{username}").Handler(getIDByUsernameHandler)
	r.Methods(http.MethodPost).Path("/user").Handler(insertUserHandler)
	r.Methods(http.MethodDelete).Path("/user").Handler(deleteUserHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+port, r))
}
