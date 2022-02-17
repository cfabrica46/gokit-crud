package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cfabrica46/gokit-crud/token-app/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load(".env") == nil {
		log.Println(".env loaded")
	}
	runServer(os.Getenv("PORT"))
}

func runServer(port string) {
	svc := service.GetService()

	err := svc.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	getGenerateTokenHandler := httptransport.NewServer(
		service.MakeGenerateTokenEndpoint(svc),
		service.DecodeGenerateTokenRequest,
		service.EncodeResponse,
	)

	getExtractTokenHandler := httptransport.NewServer(
		service.MakeExtractTokenEndpoint(svc),
		service.DecodeExtractTokenRequest,
		service.EncodeResponse,
	)

	getSetTokenHandler := httptransport.NewServer(
		service.MakeSetTokenEndpoint(svc),
		service.DecodeSetTokenRequest,
		service.EncodeResponse,
	)

	getDeleteTokenHandler := httptransport.NewServer(
		service.MakeDeleteTokenEndpoint(svc),
		service.DecodeDeleteTokenRequest,
		service.EncodeResponse,
	)

	getCheckTokenHandler := httptransport.NewServer(
		service.MakeCheckTokenEndpoint(svc),
		service.DecodeCheckTokenRequest,
		service.EncodeResponse,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path("/generate").Handler(getGenerateTokenHandler)
	r.Methods(http.MethodPost).Path("/extract").Handler(getExtractTokenHandler)
	r.Methods(http.MethodPost).Path("/token").Handler(getSetTokenHandler)
	r.Methods(http.MethodDelete).Path("/token").Handler(getDeleteTokenHandler)
	r.Methods(http.MethodPost).Path("/check").Handler(getCheckTokenHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+port, r))
}
