package service

import (
	_ "github.com/lib/pq"
)

type serviceInterface interface {
	GenerateToken(int, string, string) (string, error)
	ExtractData(string) (int, string, string, error)
}

type service struct{}

func GetService() *service {
	return &service{}
}

func (s service) GenerateToken(id int, username, email string) (token string, err error) {
	return
}

func (s service) ExtractData(token string) (id int, username, password string, err error) {
	return
}
