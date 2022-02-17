package service

import (
	_ "github.com/lib/pq"
)

type serviceInterface interface {
	SignUp(string, string, string) (string, error)
	SignIn(string, string) (string, error)
}

type service struct {
	dbHost, dbPort, tokenHost, tokenPort string
}

func GetService(dbHost, dbPort, tokenHost, tokenPort string) *service {
	return &service{dbHost, dbPort, tokenHost, tokenPort}
}

func (serviceInterface) SignUp(username, password, email string) (token string, err error) {
	return
}

func (serviceInterface) SignIn(username, password string) (token string, err error) {
	return
}
