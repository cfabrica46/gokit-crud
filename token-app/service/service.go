package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type serviceInterface interface {
	GenerateToken(int, string, string, []byte) string
	ExtractToken(string, []byte) (int, string, string, error)
	SetToken(string) error
	DeleteToken(string) error
	CheckToken(string) (bool, error)
}

//Service ...
type Service struct {
	db *redis.Client
}

//GetService ...
func GetService(db *redis.Client) *Service {
	return &Service{db}
}

//GenerateToken ...
func (Service) GenerateToken(id int, username, email string, secret []byte) (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"uuid":     uuid.NewString(),
	})

	token, _ = t.SignedString(secret)
	return
}

//ExtractToken ...
func (s Service) ExtractToken(token string, secret []byte) (id int, username, email string, err error) {
	t, err := jwt.Parse(token, keyFunc(secret))
	if err != nil {
		return
	}

	claims := t.Claims.(jwt.MapClaims)
	idAux := claims["id"].(float64)
	id = int(idAux)
	username = claims["username"].(string)
	email = claims["email"].(string)
	return
}

//SetToken ...
func (s *Service) SetToken(token string) (err error) {
	err = s.db.Set(token, true, time.Minute*10).Err()
	if err != nil {
		return
	}
	return
}

//DeleteToken ...
func (s *Service) DeleteToken(token string) error {
	return s.db.Del(token).Err()
}

//CheckToken ...
func (s Service) CheckToken(token string) (check bool, err error) {
	result, err := s.db.Get(token).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			err = nil
			return
		}
		return
	}

	if result == "1" {
		check = true
	}
	return
}

func keyFunc(secret []byte) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	}
}
