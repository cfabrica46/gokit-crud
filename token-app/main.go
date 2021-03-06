package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cfabrica46/gokit-crud/token-app/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
	}

	options := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	}
	db := redis.NewClient(options)

	runServer(os.Getenv("PORT"), db)
}

func runServer(port string, db *redis.Client) {
	svc := service.GetService(db)

	getGenerateTokenHandler := httptransport.NewServer(
		service.MakeGenerateTokenEndpoint(svc),
		service.DecodeRequest(service.IDUsernameEmailSecretRequest{}),
		service.EncodeResponse,
	)

	getExtractTokenHandler := httptransport.NewServer(
		service.MakeExtractTokenEndpoint(svc),
		service.DecodeRequest(service.TokenSecretRequest{}),
		service.EncodeResponse,
	)

	getSetTokenHandler := httptransport.NewServer(
		service.MakeManageTokenEndpoint(svc, service.NewSetTokenState()),
		service.DecodeRequest(service.Token{}),
		service.EncodeResponse,
	)

	getDeleteTokenHandler := httptransport.NewServer(
		service.MakeManageTokenEndpoint(svc, service.NewDeleteTokenState()),
		service.DecodeRequest(service.Token{}),
		service.EncodeResponse,
	)

	getCheckTokenHandler := httptransport.NewServer(
		service.MakeCheckTokenEndpoint(svc),
		service.DecodeRequest(service.Token{}),
		service.EncodeResponse,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path("/generate").Handler(getGenerateTokenHandler)
	r.Methods(http.MethodPost).Path("/extract").Handler(getExtractTokenHandler)
	r.Methods(http.MethodPost).Path("/token").Handler(getSetTokenHandler)
	r.Methods(http.MethodDelete).Path("/token").Handler(getDeleteTokenHandler)
	r.Methods(http.MethodPost).Path("/check").Handler(getCheckTokenHandler)

	log.Println("ListenAndServe on localhost:" + os.Getenv("PORT"))
	log.Println(http.ListenAndServe(":"+port, r))
}
