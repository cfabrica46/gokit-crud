package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cfabrica46/gokit-crud/database-app/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	runServer(os.Getenv("PORT"))
}

func runServer(port string) {
	svc := service.GetService()

	err := svc.OpenDB(service.DBDriver, service.PsqlInfo)
	if err != nil {
		log.Fatal(err)
	}

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
	r.Methods(http.MethodGet).Path("/user/{id:[0-9]+}").Handler(getUserByIDHandler)
	r.Methods(http.MethodGet).Path("/user/username_password").Handler(getUserByUsernameAndPasswordHandler)
	r.Methods(http.MethodGet).Path("/id/{username}").Handler(getIDByUsernameHandler)
	r.Methods(http.MethodPost).Path("/user").Handler(insertUserHandler)
	r.Methods(http.MethodDelete).Path("/user").Handler(deleteUserHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+port, r))
}
