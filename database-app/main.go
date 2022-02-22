package main

/* func main() {
	if godotenv.Load(".env") == nil {
		log.Println(".env loaded")
	}
	runServer(os.Getenv("PORT"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_SSLMODE"), os.Getenv("DB_DRIVER"))
}

func runServer(port, postgresHost, postgresPort, postgresUsername, postgresPassword, postgresDB, postgresSSL, postgresDriver string) {
	svc := service.GetService(postgresHost, postgresPort, postgresUsername, postgresPassword, postgresDB, postgresSSL, postgresDriver)

	err := svc.OpenDB()
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
} */
