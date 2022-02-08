package service

import (
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func RunServer(port string) {
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

	/* getUserByUsernameAndPasswordHandler := httptransport.NewServer(
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
	) */

	r := mux.NewRouter()
	r.Methods("GET").Path("/users").Handler(getAllUsersHandler)
	r.Methods("GET").Path("/user/{id:[0-9]+}").Handler(getUserByIDHandler)

	// http.Handle("/users", getAllUsersHandler)
	// http.Handle("/user/:id", getUserByIDHandler)
	// http.Handle("/user/usernamepassword", getUserByUsernameAndPasswordHandler)
	// http.Handle("/id/username", getIDByUsernameHandler)
	// http.Handle("/user/insert", insertUserHandler)
	// http.Handle("/user/delete", deleteUserByUsernameHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+port, r))
}
