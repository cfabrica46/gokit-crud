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

	infServ := service.InfoServices{
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		TokenHost: os.Getenv("TOKEN_HOST"),
		TokenPort: os.Getenv("TOKEN_PORT"),
		Secret:    os.Getenv("SECRET"),
	}

	runServer(
		os.Getenv("PORT"),
		&infServ,
	)
}

func runServer(port string, infServ *service.InfoServices) {
	svc := service.NewService(
		&http.Client{},
		infServ,
	)

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

	router := mux.NewRouter()
	router.Methods(http.MethodPost).Path("/signup").Handler(getSignUpHandler)
	router.Methods(http.MethodPost).Path("/signin").Handler(getSignInHandler)
	router.Methods(http.MethodPost).Path("/logout").Handler(getLogOutHandler)
	router.Methods(http.MethodGet).Path("/users").Handler(getAllUsersHandler)
	router.Methods(http.MethodPost).Path("/profile").Handler(getProfileHandler)
	router.Methods(http.MethodDelete).Path("/profile").Handler(getDeleteAccountHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Println(http.ListenAndServe(":"+port, router))
}
