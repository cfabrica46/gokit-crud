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

func (serviceInterface) SignUp(username, password, email string) (token string, err error) {
	return
}

func (serviceInterface) SignIn(username, password string) (token string, err error) {
	return
}
