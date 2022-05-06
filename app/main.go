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

type infoServices struct {
	dbHost    string
	dbPort    string
	tokenHost string
	tokenPort string
	secret    string
}

func main() {
	log.SetFlags(log.Lshortfile)

	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env loaded")
	}

	infServ := infoServices{
		dbHost:    os.Getenv("DB_HOST"),
		dbPort:    os.Getenv("DB_PORT"),
		tokenHost: os.Getenv("TOKEN_HOST"),
		tokenPort: os.Getenv("TOKEN_PORT"),
		secret:    os.Getenv("SECRET"),
	}

	runServer(
		os.Getenv("PORT"),
		&infServ,
	)
}

func runServer(port string, infServ *infoServices) {
	svc := service.NewService(
		&http.Client{},
		infServ.dbHost,
		infServ.dbHost,
		infServ.tokenHost,
		infServ.tokenPort,
		infServ.secret,
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
