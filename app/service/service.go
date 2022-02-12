package service

import (
	_ "github.com/lib/pq"
)

type serviceInterface interface {
	SignUp(string, string, string) (string, error)
	SignIn(string, string) (string, error)
}

type service struct{}

func GetService() *service {
	return &service{}
}

func (s service) GetAllUsers(username, password string) (token string, err error) {
	return
}
