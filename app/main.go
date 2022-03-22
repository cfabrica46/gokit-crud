package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cfabrica46/gokit-crud/app/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Lshortfile)
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env loaded")
	}

	runServer(os.Getenv("PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("TOKEN_HOST"), os.Getenv("TOKEN_PORT"), os.Getenv("SECRET"))
}

func runServer(port, dbHost, dbPort, tokenHost, tokenPort, secret string) {
	svc := service.NewService(&http.Client{}, dbHost, dbPort, tokenHost, tokenPort, secret)

	getSignUpHandler := httptransport.NewServer(
		service.MakeSignUpEndpoint(svc),
		service.DecodeSignUpRequest,
		service.EncodeResponse,
	)

	getSignInHandler := httptransport.NewServer(
		service.MakeSignInEndpoint(svc),
		service.DecodeSignInRequest,
		service.EncodeResponse,
	)

	getLogOutHandler := httptransport.NewServer(
		service.MakeLogOutEndpoint(svc),
		service.DecodeLogOutRequest,
		service.EncodeResponse,
	)

	getAllUsersHandler := httptransport.NewServer(
		service.MakeGetAllUsersEndpoint(svc),
		service.DecodeGetAllUsersRequest,
		service.EncodeResponse,
	)

	getProfileHandler := httptransport.NewServer(
		service.MakeProfileEndpoint(svc),
		service.DecodeProfileRequest,
		service.EncodeResponse,
	)

	getDeleteAccountHandler := httptransport.NewServer(
		service.MakeDeleteAccountEndpoint(svc),
		service.DecodeDeleteAccountRequest,
		service.EncodeResponse,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path("/signup").Handler(getSignUpHandler)
	r.Methods(http.MethodPost).Path("/signin").Handler(getSignInHandler)
	r.Methods(http.MethodPost).Path("/logout").Handler(getLogOutHandler)
	r.Methods(http.MethodGet).Path("/users").Handler(getAllUsersHandler)
	r.Methods(http.MethodPost).Path("/profile").Handler(getProfileHandler)
	r.Methods(http.MethodDelete).Path("/profile").Handler(getDeleteAccountHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+port, r))
}
