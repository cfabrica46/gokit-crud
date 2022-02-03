package service

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

func runServer() {
	svc := serviceDB{}

	/* GetAllUsers() ([]models.User, error)
	   GetUserByID(int) (models.User, error)
	   GetUserByUsernameAndPassword(string, string) (models.User, error)
	   GetIDByUsername(string) (int, error)
	   InsertUser(string, string, string) error
	   DeleteUserByUsername(string) error */

	getAllUsersHandler := httptransport.NewServer(
		makeGetAllUsersEndpoint(svc),
		decodeGetAllUsersRequest,
		encodeResponse,
	)
}
